package bridge

import (
	"store/internal/transaction"
	"time"
)

type Date interface {
	Today() time.Time
	Now() time.Time
}

type Notification interface {
	SendUser(ts transaction.Session, userID int, title, message string) error
	SendAll(ts transaction.Session, title, message string) error
}
