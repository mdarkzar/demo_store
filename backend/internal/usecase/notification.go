package usecase

import (
	"fmt"
	"store/internal/entity/global"
	"store/internal/entity/notification"
	"store/internal/rimport"
	"store/internal/transaction"

	"github.com/sirupsen/logrus"
)

type NotificationUsecase struct {
	log *logrus.Logger
	rimport.RepositoryImports
}

func NewNotificationUsecase(
	log *logrus.Logger,
	ri rimport.RepositoryImports,
) *NotificationUsecase {
	return &NotificationUsecase{
		log:               log,
		RepositoryImports: ri,
	}
}

func (u *NotificationUsecase) logPrefix() string {
	return "[notification_usecase]"
}

func (u *NotificationUsecase) SendUser(ts transaction.Session, userID int, title, message string) error {
	lf := logrus.Fields{
		"user_id": userID,
	}

	messageID, err := u.Repository.Notification.Create(ts, title, message)
	if err != nil {
		u.log.WithFields(lf).Errorln(u.logPrefix(),
			fmt.Sprintf("не удалось создать сообщение; ошибка: %v", err),
		)
		return global.ErrInternalError
	}

	err = u.Repository.Notification.CreateUserMessage(ts, userID, messageID)
	if err != nil {
		u.log.WithFields(lf).Errorln(u.logPrefix(),
			fmt.Sprintf("не удалось создать сообщение для пользователя; ошибка: %v", err),
		)
		return global.ErrInternalError
	}

	u.log.WithFields(lf).Debugln(u.logPrefix(), "создано сообщение id", messageID)

	return nil
}

func (u *NotificationUsecase) SendAll(ts transaction.Session, title, message string) error {
	lf := logrus.Fields{
		"title": title,
	}

	userList, err := u.Repository.User.LoadAll(ts)
	if err != nil {
		u.log.WithFields(lf).Errorln(u.logPrefix(),
			fmt.Sprintf("не удалось загрузить пользователей; ошибка: %v", err),
		)
		return global.ErrInternalError
	}

	messageID, err := u.Repository.Notification.Create(ts, title, message)
	if err != nil {
		u.log.WithFields(lf).Errorln(u.logPrefix(),
			fmt.Sprintf("не удалось создать сообщение; ошибка: %v", err),
		)
		return global.ErrInternalError
	}

	for _, row := range userList {
		err := u.Repository.Notification.CreateUserMessage(ts, row.ID, messageID)
		if err != nil {
			u.log.WithFields(lf).Errorln(u.logPrefix(),
				fmt.Sprintf("не удалось создать пользовательское сообщение; ошибка: %v", err),
			)
			continue
		}

		lf["userID"] = row.ID

		u.log.WithFields(lf).Debugln(u.logPrefix(), "создано сообщение id", messageID)
	}

	return nil
}

func (u *NotificationUsecase) LoadUserMessages(ts transaction.Session, userID int) ([]notification.Notification, error) {
	lf := logrus.Fields{
		"user_id": userID,
	}

	messageList, err := u.Repository.Notification.FindUserMessages(ts, userID)
	switch err {
	case global.ErrNoData:
		u.log.WithFields(lf).Debugln(u.logPrefix(), "нет новых сообщений")
		return nil, nil
	case nil:
		for _, row := range messageList {
			if err := u.Repository.Notification.Delete(ts, row.ID, userID); err != nil {
				u.log.WithFields(lf).Errorln(u.logPrefix(),
					fmt.Sprintf("ошибка при удаления сообщения; ошибка: %v", err),
				)
			}
		}

		return messageList, nil

	default:
		u.log.WithFields(lf).Errorln(u.logPrefix(),
			fmt.Sprintf("не удалось получить сообщения; ошибка: %v", err),
		)
		return nil, global.ErrInternalError
	}

}
