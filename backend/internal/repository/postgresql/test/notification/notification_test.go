package notification_test

import (
	"os"
	"store/config"
	"store/internal/entity/global"
	"store/internal/repository/postgresql"
	"store/internal/transaction"
	"store/tools/pgdb"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	r := require.New(t)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)
	r.NotEmpty(conf)

	db := pgdb.SqlxDB(conf.PostgresURL())
	r.NoError(db.Ping())

	ts := transaction.NewSQLSession(db)
	ts.Start()
	defer ts.Rollback()

	repo := postgresql.NewNotification()

	var (
		title   string
		message string
	)

	faker.FakeData(&title)
	faker.FakeData(&message)

	messageID, err := repo.Create(ts, title, message)
	r.NoError(err)
	r.NotEmpty(messageID)
}

func TestCreateUserMessage(t *testing.T) {
	r := require.New(t)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)
	r.NotEmpty(conf)

	db := pgdb.SqlxDB(conf.PostgresURL())
	r.NoError(db.Ping())

	ts := transaction.NewSQLSession(db)
	ts.Start()
	defer ts.Rollback()

	repo := postgresql.NewNotification()

	var (
		userID  = 0
		title   string
		message string
	)

	faker.FakeData(&title)
	faker.FakeData(&message)

	messageID, err := repo.Create(ts, title, message)
	r.NoError(err)
	r.NotEmpty(messageID)

	err = repo.CreateUserMessage(ts, userID, messageID)
	r.NoError(err)
}

func TestFindUserMessages(t *testing.T) {
	r := require.New(t)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)
	r.NotEmpty(conf)

	db := pgdb.SqlxDB(conf.PostgresURL())
	r.NoError(db.Ping())

	ts := transaction.NewSQLSession(db)
	ts.Start()
	defer ts.Rollback()

	repo := postgresql.NewNotification()

	var (
		userID  = 0
		title   string
		message string
	)

	faker.FakeData(&title)
	faker.FakeData(&message)

	messageID, err := repo.Create(ts, title, message)
	r.NoError(err)
	r.NotEmpty(messageID)

	err = repo.CreateUserMessage(ts, userID, messageID)
	r.NoError(err)

	messageList, err := repo.FindUserMessages(ts, userID)
	r.NoError(err)
	r.NotEmpty(messageList)
}

func TestDelete(t *testing.T) {
	r := require.New(t)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)
	r.NotEmpty(conf)

	db := pgdb.SqlxDB(conf.PostgresURL())
	r.NoError(db.Ping())

	ts := transaction.NewSQLSession(db)
	ts.Start()
	defer ts.Rollback()

	repo := postgresql.NewNotification()

	var (
		userID  = 0
		title   string
		message string
	)

	faker.FakeData(&title)
	faker.FakeData(&message)

	_, err = postgresql.SqlxTx(ts).Exec(`DELETE FROM user$notification`)
	r.NoError(err)

	messageID, err := repo.Create(ts, title, message)
	r.NoError(err)
	r.NotEmpty(messageID)

	err = repo.CreateUserMessage(ts, userID, messageID)
	r.NoError(err)

	err = repo.Delete(ts, messageID, userID)
	r.NoError(err)

	messageList, err := repo.FindUserMessages(ts, userID)
	r.Error(err)
	r.Equal(global.ErrNoData, err)
	r.Empty(messageList)
}
