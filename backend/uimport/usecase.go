package uimport

import "store/internal/usecase"

type Usecase struct {
	User         *usecase.UserUsecase
	Product      *usecase.ProductUsecase
	Notification *usecase.NotificationUsecase
	Queue        *usecase.QueueUsecase
}
