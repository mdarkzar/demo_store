package rabbitmq

import (
	"store/config"
	"store/internal/entity/queue"
	"store/internal/repository"

	"github.com/streadway/amqp"
)

type queueRepository struct {
	config config.Config
	// rabbitmq
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewQueue(config config.Config) repository.Queue {
	return &queueRepository{
		config: config,
	}
}

func (r *queueRepository) Connect() (err error) {
	r.conn, err = amqp.Dial(r.config.RabbitMQConnectURL())
	if err != nil {
		err = queue.NewQueueError(err)
		return
	}

	r.ch, err = r.conn.Channel()
	if err != nil {
		err = queue.NewQueueError(err)
		return
	}

	return
}

func (r *queueRepository) QueueDeclare(routingKey string) error {
	_, err := r.ch.QueueDeclare(
		routingKey, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return queue.NewQueueError(err)
	}

	return nil
}

func (r *queueRepository) QueueBind(queueName, routingKey string, exchange string) error {
	if err := r.ch.QueueBind(
		queueName,  // name of the queue
		routingKey, // bindingKey
		exchange,   // sourceExchange
		false,      // noWait
		nil,        // arguments
	); err != nil {
		return queue.NewQueueError(err)
	}

	return nil
}

func (r *queueRepository) ExchangeDeclare(name, exchangeType string) error {
	if err := r.ch.ExchangeDeclare(
		name,         // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return queue.NewQueueError(err)
	}

	return nil
}

func (r *queueRepository) Disconnect() (err error) {
	if r.conn != nil {
		if err = r.conn.Close(); err != nil {
			err = queue.NewQueueError(err)
			return
		}
	}

	if r.ch != nil {
		if err = r.ch.Close(); err != nil {
			err = queue.NewQueueError(err)
			return
		}
	}

	return nil
}

func (r *queueRepository) Write(routingKey, exchange string, message []byte, contentType queue.MessageType) (err error) {
	if err = r.ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  string(contentType),
			Body:         message,
			DeliveryMode: 2,
		},
	); err != nil {
		err = queue.NewQueueError(err)
		return
	}

	return
}

func (r *queueRepository) ListenQueue(queueName string, consumerTag string) (data <-chan amqp.Delivery, err error) {
	data, err = r.ch.Consume(
		queueName,   // queue
		consumerTag, // consumer
		false,       // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		err = queue.NewQueueError(err)
		return
	}

	return
}

func (r *queueRepository) IsClosed() bool {
	return r.conn == nil || r.conn.IsClosed()
}
