package usecase

import (
	"context"
	"fmt"
	"store/internal/entity/queue"
	"store/internal/proto/notification/notificationproto"
	"time"

	"google.golang.org/protobuf/proto"
)

func (u *NotificationUsecase) ListenNotifications(termFlag chan struct{}) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go u.Bridge.Queue.ConnectionControl(ctx)
	u.Bridge.Queue.WaitConnectionInitialized()

	if err := u.Bridge.Queue.DeclareQueue(queue.NotificationQueue); err != nil {
		return
	}

	for {
		if err := u.listenNotification(ctx, termFlag); err == nil {
			break
		} else if err == queue.ErrQueueUnavailable {
			u.log.Errorln("отсутствует связь с очередью, переподключиться через", queue.QueueReconnectInterval)
			time.Sleep(queue.QueueReconnectInterval)
		}
	}
}

func (u *NotificationUsecase) listenNotification(ctx context.Context, termFlag chan struct{}) error {
	queueName := queue.NotificationQueue
	taskChannel, err := u.Bridge.Queue.Listen(queueName, "")
	if err != nil {
		u.log.Errorln(fmt.Sprintf("не удалось слушать с очереди %s; ошибка: %v", queueName, err))
		return queue.ErrQueueUnavailable
	}

Loop:
	for {
		select {
		case <-termFlag:
			u.log.Warnf("** аварийная остановка %s, ждите завершения **", queueName)
			break Loop
		case m, ok := <-taskChannel:
			if !ok {
				return queue.ErrQueueUnavailable
			}

			func() {
				defer m.Ack(true)

				ts := u.SessionManager.CreateSession()
				if err := ts.Start(); err != nil {
					u.log.Errorln(fmt.Sprintf("не удалось открыть транзакцию; ошибка: %v", err))
					return
				}
				defer ts.Rollback()

				var notif notificationproto.Notification

				if err = proto.Unmarshal(m.Body, &notif); err != nil {
					u.log.Errorln(fmt.Sprintf("не удалось пропарсить json %s; ошибка: %v", string(m.Body), err))
					time.Sleep(queue.QueueErrorInterval)
					return
				}

				u.SendAll(ts, notif.Title, notif.Message)
				u.log.Debugln("обработана отправка сообщения")

				if err = ts.Commit(); err != nil {
					u.log.Errorln("не удалось закрыть транзакцию; ошибка:", err)
				}
			}()
		}
	}

	return nil
}
