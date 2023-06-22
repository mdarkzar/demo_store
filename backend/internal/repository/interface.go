package repository

import (
	"store/internal/entity/notification"
	"store/internal/entity/product"
	"store/internal/entity/queue"
	"store/internal/entity/user"
	"store/internal/transaction"

	"github.com/streadway/amqp"
)

type Product interface {
	Create(ts transaction.Session, userID int, name string, price float64, stID int) (productID int, err error)
	Remove(ts transaction.Session, productID int) error
	FindByID(ts transaction.Session, productID int) (product.Product, error)
	LoadAll(ts transaction.Session) ([]product.Product, error)
	LoadStorageList(ts transaction.Session) ([]product.Storage, error)
}

type User interface {
	Create(ts transaction.Session, login, password string) (userID int, err error)
	FindByLogin(ts transaction.Session, login string) (user.User, error)
	FindByID(ts transaction.Session, userID int) (user.User, error)
	LoadAll(ts transaction.Session) ([]user.User, error)
}

type UserCache interface {
	Add(userID int, user user.User)
	Remove(userID int)
	Get(userID int) (user user.User, exists bool)
}

type Notification interface {
	Create(ts transaction.Session, title string, message string) (messageID int, err error)
	CreateUserMessage(ts transaction.Session, userID, messageID int) error
	FindUserMessages(ts transaction.Session, userID int) ([]notification.Notification, error)
	Delete(ts transaction.Session, nID, userID int) error
}

type Queue interface {
	Connect() error
	Disconnect() error
	Write(routingKey, exchange string, message []byte, contentType queue.MessageType) (err error)
	QueueDeclare(routingKey string) error
	QueueBind(queueName, routingKey string, exchange string) error
	ExchangeDeclare(name, exchangeType string) error
	ListenQueue(queueName string, consumerTag string) (data <-chan amqp.Delivery, err error)
	IsClosed() bool
}
