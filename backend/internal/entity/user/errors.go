package user

import "errors"

var (
	ErrLoginOrPasswordIncorrect = errors.New("некорректный логин или пароль")
)
