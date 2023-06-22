package usecase

import (
	"context"
	"fmt"
	"store/internal/entity/global"
	"store/internal/entity/queue"
	"store/internal/rimport"

	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type QueueUsecase struct {
	// вначале системные объекты - логи, конфиги
	log *logrus.Logger
	// далее репозитории
	rimport.RepositoryImports
	// приватные значения
	reconnectCheckInterval time.Duration
	firstConnectLock       chan struct{}
	init                   bool
}

func NewQueue(
	log *logrus.Logger,
	ri rimport.RepositoryImports,
) *QueueUsecase {
	return &QueueUsecase{
		log:                    log,
		RepositoryImports:      ri,
		reconnectCheckInterval: queue.ReconnectInterval,
		firstConnectLock:       make(chan struct{}),
	}
}

// ConnectionControl подключение и контроль подключения
func (u *QueueUsecase) ConnectionControl(ctx context.Context) {
	if u.reconnectCheckInterval == 0 {
		u.log.Fatalln("не задан reconnectCheckInterval")
	}

	u.establishConnection()
	u.log.Debugln("запущен контроль над подключением к очереди, время проверки каждые", u.reconnectCheckInterval)

	checkTicker := time.NewTicker(u.reconnectCheckInterval)
	for {
		select {
		case <-checkTicker.C:
			if u.Repository.Queue.IsClosed() {
				u.log.Warnln("отсутствие соединение с очередью, подключение")
				u.establishConnection()
				u.log.Infoln("успешное подключение к очереди")
			}
		case <-ctx.Done():
			u.Repository.Queue.Disconnect()
			return
		}
	}
}

// establishConnection установка соединения и инициализации очереди
func (u *QueueUsecase) establishConnection() (err error) {
	if u.Repository.Queue.IsClosed() {
		u.Repository.Queue.Disconnect()
		if err = u.Repository.Queue.Connect(); err != nil {
			u.log.Errorln(fmt.Sprintf("не удалось подключиться; ошибка: %v", err))
			return global.ErrInternalError
		}

		if !u.init {
			close(u.firstConnectLock)
			u.init = true
		}

		u.log.Debugln("соединение установлено")
	}

	return
}

// Listen прослушивание очереди
func (u *QueueUsecase) Listen(routingKey, consumerTag string) (taskChannel <-chan amqp.Delivery, err error) {
	u.log.Debugln(fmt.Sprintf("начато прослушивание очереди %s", routingKey))

	return u.Repository.Queue.ListenQueue(routingKey, consumerTag)
}

// Write запись в очередь
func (u *QueueUsecase) Write(routingKey, exchange string, message []byte, contentType queue.MessageType) error {
	u.log.Debugln(fmt.Sprintf("отправка таска в очередь %s", routingKey))

	lf := logrus.Fields{
		"queue":   routingKey,
		"message": string(message),
	}

	if err := u.Repository.Queue.Write(routingKey, exchange, message, queue.Json); err != nil {
		u.log.WithFields(lf).Errorln(
			fmt.Sprintf("не удалось отправить таск в очередь %s; ошибка: %v",
				routingKey, err))
		return global.ErrInternalError
	}

	return nil
}

// DeclareQueue объявление очереди
func (u *QueueUsecase) DeclareQueue(queueName string) error {
	u.WaitConnectionInitialized() // необходимо, чтобы код выполнялся после первого подключения в горутине
	err := u.Repository.Queue.QueueDeclare(queueName)
	if err != nil {
		u.log.Errorln(fmt.Sprintf("не удалось создать очередь; ошибка: %v", err))
		return global.ErrInternalError
	}

	u.log.Debugln("очередь создана:", queueName)

	return nil
}

// ExchangeDeclare объявление exchange
func (u *QueueUsecase) ExchangeDeclare(exchange, exchangeType string) error {
	u.WaitConnectionInitialized() // необходимо, чтобы код выполнялся после первого подключения в горутине
	err := u.Repository.Queue.ExchangeDeclare(exchange, exchangeType)
	if err != nil {
		u.log.Errorln(fmt.Sprintf("не удалось создать exchange; ошибка: %v", err))
		return global.ErrInternalError
	}

	u.log.Debugln("exchange создан:", exchange)

	return nil
}

// QueueDeclareWithBind создание очереди и её привязка к exchange
func (u *QueueUsecase) QueueDeclareWithBind(queueName, routingKey, exchangeName string) error {
	u.WaitConnectionInitialized() // необходимо, чтобы код выполнялся после первого подключения в горутине
	err := u.Repository.Queue.QueueDeclare(queueName)
	if err != nil {
		u.log.Errorln(fmt.Sprintf("не удалось создать очередь; ошибка: %v", err))
		return global.ErrInternalError
	}

	err = u.Repository.Queue.QueueBind(queueName, routingKey, exchangeName)
	if err != nil {
		u.log.Errorln(fmt.Sprintf("не удалось привязать очередь %s; ошибка: %v", queueName, err))
		return global.ErrInternalError
	}

	u.log.WithFields(logrus.Fields{
		"queue":    queueName,
		"key":      routingKey,
		"exchange": exchangeName,
	}).Debugln("очередь создана:", queueName, "и привязана:", exchangeName)

	return nil
}

func (u *QueueUsecase) WaitConnectionInitialized() {
	<-u.firstConnectLock
}
