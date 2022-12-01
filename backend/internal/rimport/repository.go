package rimport

import "store/internal/repository"

type Repository struct {
	Product      repository.Product
	User         repository.User
	UserCache    repository.UserCache
	Notification repository.Notification
}

type MockRepository struct {
	Product      *repository.MockProduct
	User         *repository.MockUser
	UserCache    *repository.MockUserCache
	Notification *repository.MockNotification
}
