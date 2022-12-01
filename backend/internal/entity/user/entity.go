package user

import "time"

type User struct {
	ID          int       `json:"id" db:"user_id"`
	Login       string    `json:"login" db:"login"`
	Password    string    `json:"-" db:"password"`
	CreatedDate time.Time `json:"-" db:"created_date"`
}
