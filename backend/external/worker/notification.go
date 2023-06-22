package worker

import (
	"store/tools/ossignal"
	"store/uimport"

	"github.com/sirupsen/logrus"
)

type NotificationWorker struct {
	log *logrus.Logger
	uimport.UsecaseImports
}

func NewNotificationWorker(ui uimport.UsecaseImports,
	log *logrus.Logger,
) *NotificationWorker {
	return &NotificationWorker{
		log:            log,
		UsecaseImports: ui,
	}
}

func (e *NotificationWorker) Run() {
	termFlag := make(chan struct{})
	go ossignal.WaitForTerm(termFlag)

	e.Usecase.Notification.ListenNotifications(termFlag)
}
