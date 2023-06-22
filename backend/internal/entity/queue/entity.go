package queue

import (
	"encoding/json"
	"time"

	"github.com/streadway/amqp"
)

type Task struct {
	m amqp.Delivery
}

func (t *Task) Ready() {
	t.m.Ack(true)
}

func NewQueueData(m amqp.Delivery) Task {
	return Task{
		m: m,
	}
}

func (t *Task) Unmarshal(obj interface{}) error {
	return json.Unmarshal(t.m.Body, obj)
}

type TestData struct {
	String string
	Number int
	Date   time.Time
}
