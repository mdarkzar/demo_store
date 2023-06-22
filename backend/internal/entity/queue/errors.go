package queue

import (
	"errors"
	"fmt"
)

type QueueError struct {
	err error
}

func (q *QueueError) Error() string {
	return fmt.Sprintf("внутреняя ошибка очереди: %v", q.err.Error())
}

func NewQueueError(err error) error {
	return &QueueError{err: err}
}

var (
	ErrQueueUnavailable = errors.New("очередь недоступна")
)
