package product_test

import (
	"os"
	"store/config"
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

	repo := postgresql.NewUser()

	var (
		userID   = 0
		login    string
		password string
	)

	faker.FakeData(&login)
	faker.FakeData(&password)

	userID, err = repo.Create(ts, login, password)
	r.NoError(err)
	r.NotEmpty(userID)

}

func TestFindByID(t *testing.T) {
	r := require.New(t)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)
	r.NotEmpty(conf)

	db := pgdb.SqlxDB(conf.PostgresURL())
	r.NoError(db.Ping())

	ts := transaction.NewSQLSession(db)
	ts.Start()
	defer ts.Rollback()

	repo := postgresql.NewUser()

	var (
		userID   = 0
		login    string
		password string
	)

	faker.FakeData(&login)
	faker.FakeData(&password)

	userID, err = repo.Create(ts, login, password)
	r.NoError(err)
	r.NotEmpty(userID)

	userData, err := repo.FindByID(ts, userID)
	r.NoError(err)
	r.Equal(login, userData.Login)
	r.Equal(password, userData.Password)
	r.Equal(userID, userData.ID)

}

func TestLoadAll(t *testing.T) {
	r := require.New(t)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)
	r.NotEmpty(conf)

	db := pgdb.SqlxDB(conf.PostgresURL())
	r.NoError(db.Ping())

	ts := transaction.NewSQLSession(db)
	ts.Start()
	defer ts.Rollback()

	repo := postgresql.NewUser()

	var (
		userID   = 0
		login    string
		password string
	)

	faker.FakeData(&login)
	faker.FakeData(&password)

	userID, err = repo.Create(ts, login, password)
	r.NoError(err)
	r.NotEmpty(userID)

	userList, err := repo.LoadAll(ts)
	r.NoError(err)
	r.NotEmpty(userList)

	founded := false
	for _, row := range userList {
		if row.ID == userID {
			founded = true
			break
		}
	}
	r.True(founded)

}
