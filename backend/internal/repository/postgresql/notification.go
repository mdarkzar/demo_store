package postgresql

import (
	"store/internal/entity/notification"
	"store/internal/repository"
	"store/internal/transaction"
	"store/tools/gensql"
)

type notificationRepo struct{}

func NewNotification() repository.Notification {
	return &notificationRepo{}
}

func (r *notificationRepo) Create(ts transaction.Session, title string, message string) (messageID int, err error) {
	return gensql.Get[int](SqlxTx(ts), `INSERT INTO notification ( title, message) VALUES ($1, $2) returning n_id`, title, message)
}

func (r *notificationRepo) CreateUserMessage(ts transaction.Session, userID, messageID int) error {
	sqlQuery := `INSERT INTO user$notification ( n_id, user_id) VALUES ($1, $2)`
	_, err := SqlxTx(ts).Exec(sqlQuery, messageID, userID)

	return err
}

func (r *notificationRepo) FindUserMessages(ts transaction.Session, userID int) ([]notification.Notification, error) {
	sqlQuery := `
	select n.n_id, un.user_id, n.title, n.message, n.created_date
	from notification n
	       join public.user$notification un on (un.n_id = n.n_id)
	where un.user_id = $1
	order by created_date desc
	`

	return gensql.Select[notification.Notification](SqlxTx(ts), sqlQuery, userID)
}

func (r *notificationRepo) Delete(ts transaction.Session, nID, userID int) error {
	_, err := SqlxTx(ts).Exec(`DELETE FROM user$notification WHERE user_id = $1 and n_id = $2`, userID, nID)

	return err
}
