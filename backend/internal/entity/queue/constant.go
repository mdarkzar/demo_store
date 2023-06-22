package queue

import "time"

const (
	NotificationQueue = "NOTIFICATION_QUEUE"
)

type MessageType string

const (
	Json  MessageType = "application/json"
	Proto MessageType = "application/protobuf"
)

const (
	// ExchangeTypeTopic в режиме topic маршрутизация осуществляется по routing key https://www.youtube.com/watch?v=PNNCiy4DseA&t=111s
	ExchangeTypeTopic = "topic"
)

const (
	EmptyExchange = ""
)

const (
	// ReconnectInterval переподключение к очереди
	ReconnectInterval = time.Second * 5
	// QueueErrorInterval интервал при проблеме с очередью
	QueueErrorInterval = time.Second * 1
	// QueueReconnectInterval переподключение к очереди
	QueueReconnectInterval = time.Second * 5
	// ErrorExtendInterval интервал, который выставляется в случае, когда невозможно отправить сообщение
	ErrorExtendInterval = time.Minute * 5
)
