package product_test

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

	repo := postgresql.NewProduct()

	var (
		userID = 0
		name   string
		price  float64
	)

	faker.FakeData(&name)
	faker.FakeData(&price)

	productID, err := repo.Create(ts, userID, name, price)
	r.NoError(err)
	r.NotEmpty(productID)

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

	repo := postgresql.NewProduct()

	var (
		userID = 0
		name   string
		price  float64
	)

	faker.FakeData(&name)
	faker.FakeData(&price)

	productID, err := repo.Create(ts, userID, name, price)
	r.NoError(err)
	r.NotEmpty(productID)

	productData, err := repo.FindByID(ts, productID)
	r.NoError(err)
	r.Equal(name, productData.Name)
	r.Equal(price, productData.Price)
	r.Equal(productID, productData.ID)

}

func TestRemove(t *testing.T) {
	r := require.New(t)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)
	r.NotEmpty(conf)

	db := pgdb.SqlxDB(conf.PostgresURL())
	r.NoError(db.Ping())

	ts := transaction.NewSQLSession(db)
	ts.Start()
	defer ts.Rollback()

	repo := postgresql.NewProduct()

	var (
		userID = 0
		name   string
		price  float64
	)

	faker.FakeData(&name)
	faker.FakeData(&price)

	productID, err := repo.Create(ts, userID, name, price)
	r.NoError(err)
	r.NotEmpty(productID)

	err = repo.Remove(ts, productID)
	r.NoError(err)

	_, err = repo.FindByID(ts, productID)
	r.Error(err)
	r.Equal(global.ErrNoData, err)

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

	repo := postgresql.NewProduct()

	var (
		userID = 0
		name   string
		price  float64
	)

	faker.FakeData(&name)
	faker.FakeData(&price)

	productID, err := repo.Create(ts, userID, name, price)
	r.NoError(err)
	r.NotEmpty(productID)

	productAll, err := repo.LoadAll(ts)
	r.NoError(err)
	r.NotEmpty(productAll)

	founded := false
	for _, row := range productAll {
		if row.ID == productID {
			founded = true
			break
		}
	}
	r.True(founded)

}
