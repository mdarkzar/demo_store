package notification

import "time"

type Notification struct {
	ID          int       `json:"-" db:"n_id"`
	UserID      int       `json:"-" db:"user_id"`
	Title       string    `json:"title" db:"title"`
	Message     string    `json:"message" db:"message"`
	CreatedDate time.Time `json:"created_date" db:"created_date"`
}
