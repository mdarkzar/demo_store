package bridge

import (
	"context"
	"store/internal/entity/queue"
	"store/internal/transaction"
	"time"

	"github.com/streadway/amqp"
)

type Date interface {
	Today() time.Time
	Now() time.Time
}

type Notification interface {
	SendUser(ts transaction.Session, userID int, title, message string) error
	SendAll(ts transaction.Session, title, message string) error
	SendAllViaQueue(title, message string) error
}

type Queue interface {
	ConnectionControl(ctx context.Context)
	Listen(routingKey, consumerTag string) (taskChannel <-chan amqp.Delivery, err error)
	Write(routingKey, exchange string, message []byte, contentType queue.MessageType) (err error)
	DeclareQueue(queueName string) error
	ExchangeDeclare(exchange, exchangeType string) error
	QueueDeclareWithBind(queueName, routingKey, exchangeName string) error
	WaitConnectionInitialized()
}
