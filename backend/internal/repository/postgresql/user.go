package postgresql

import (
	"store/internal/entity/user"
	"store/internal/repository"
	"store/internal/transaction"
	"store/tools/gensql"
)

type userRepo struct{}

func NewUser() repository.User {
	return &userRepo{}
}

func (r *userRepo) Create(ts transaction.Session, login string, password string) (userID int, err error) {
	return gensql.Get[int](SqlxTx(ts), `INSERT INTO users (login, password) VALUES ($1, $2) returning user_id`, login, password)
}

func (r *userRepo) FindByLogin(ts transaction.Session, login string) (user.User, error) {
	sqlQuery := `
	select u.user_id, u.login, u.password, u.created_date
	from users u 
	where u.login = $1
	`

	return gensql.Get[user.User](SqlxTx(ts), sqlQuery, login)
}

func (r *userRepo) FindByID(ts transaction.Session, userID int) (user.User, error) {
	sqlQuery := `
	select u.user_id, u.login, u.password, u.created_date
	from users u 
	where u.user_id = $1
	`

	return gensql.Get[user.User](SqlxTx(ts), sqlQuery, userID)
}

func (r *userRepo) LoadAll(ts transaction.Session) ([]user.User, error) {
	sqlQuery := `
	select u.user_id, u.login, u.password, u.created_date
	from users u 
	order by u.created_date
	`

	return gensql.Select[user.User](SqlxTx(ts), sqlQuery)
}
