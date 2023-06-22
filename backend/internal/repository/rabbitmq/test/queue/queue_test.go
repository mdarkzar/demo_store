package queue_test

import (
	"encoding/json"
	"store/config"
	"store/internal/entity/queue"
	"store/internal/repository/rabbitmq"

	"os"
	"sync"

	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConnect(t *testing.T) {
	r := require.New(t)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)
	r.NotEmpty(conf)

	repo := rabbitmq.NewQueue(conf)
	err = repo.Connect()
	defer repo.Disconnect()

	r.NoError(err)

}

func TestDisconnect(t *testing.T) {
	r := require.New(t)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)
	r.NotEmpty(conf)

	repo := rabbitmq.NewQueue(conf)
	err = repo.Connect()
	r.NoError(err)
	defer repo.Disconnect()

	err = repo.Disconnect()
	r.NoError(err)
}

func TestDeclareQueue(t *testing.T) {
	r := require.New(t)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)
	r.NotEmpty(conf)

	repo := rabbitmq.NewQueue(conf)
	err = repo.Connect()
	defer repo.Disconnect()

	r.NoError(err)

	err = repo.QueueDeclare("test")
	r.NoError(err)
}

func TestWriteListenQueue(t *testing.T) {
	r := require.New(t)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)
	r.NotEmpty(conf)

	repo := rabbitmq.NewQueue(conf)
	err = repo.Connect()
	r.NoError(err)

	defer repo.Disconnect()

	date := time.Now()

	var (
		data = queue.TestData{
			String: "test",
			Number: 1,
			Date:   date,
		}
		routingKey = "test"
	)

	err = repo.QueueDeclare(routingKey)
	r.NoError(err)

	jsonData, err := json.Marshal(data)
	r.NoError(err)

	err = repo.Write(routingKey, queue.EmptyExchange, jsonData, queue.Json)
	r.NoError(err)

	listenQueue, err := repo.ListenQueue(routingKey, "test")
	r.NoError(err)

	for m := range listenQueue {

		task := queue.NewQueueData(m)
		task.Ready()

		var dataRow queue.TestData

		err = task.Unmarshal(&dataRow)
		r.NoError(err)

		r.Equal(data.Number, dataRow.Number)
		r.Equal(data.String, dataRow.String)
		break
	}
}

func TestQueueWithDisconnect(t *testing.T) {
	r := require.New(t)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)
	r.NotEmpty(conf)

	repo := rabbitmq.NewQueue(conf)
	err = repo.Connect()
	r.NoError(err)

	err = repo.Connect()
	r.NoError(err)

	defer repo.Disconnect()

	date := time.Now()

	var (
		data = queue.TestData{
			String: "test",
			Number: 1,
			Date:   date,
		}
		routingKey = "test"
	)

	err = repo.QueueDeclare(routingKey)
	r.NoError(err)

	jsonData, err := json.Marshal(data)
	r.NoError(err)

	err = repo.Write(routingKey, queue.EmptyExchange, jsonData, queue.Json)
	r.NoError(err)

	listenQueue, err := repo.ListenQueue(routingKey, "test")
	r.NoError(err)

	err = repo.Disconnect()
	r.NoError(err)

	for {
		m, ok := <-listenQueue
		if !ok {

			for {
				if err = repo.Connect(); err != nil {
					time.Sleep(5 * time.Second)
					continue
				}

				break
			}

			listenQueue, err = repo.ListenQueue(routingKey, "test")
			r.NoError(err)
			continue
		}

		task := queue.NewQueueData(m)

		var dataRow queue.TestData

		err = task.Unmarshal(&dataRow)
		r.NoError(err)

		r.Equal(data.Number, dataRow.Number)
		r.Equal(data.String, dataRow.String)
		break
	}
}

func TestIsClosed(t *testing.T) {
	r := require.New(t)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)
	r.NotEmpty(conf)

	repo := rabbitmq.NewQueue(conf)
	err = repo.Connect()

	r.NoError(err)

	r.False(repo.IsClosed())

	repo.Disconnect()

	r.True(repo.IsClosed())
}

func TestWriteListenExchangeQueue(t *testing.T) {
	r := require.New(t)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NoError(err)
	r.NotEmpty(conf)

	repo := rabbitmq.NewQueue(conf)
	err = repo.Connect()
	r.NoError(err)

	defer repo.Disconnect()

	date := time.Now()

	var (
		data = queue.TestData{
			String: "test",
			Number: 1,
			Date:   date,
		}
		routingKey  = "test.#"
		routingKey2 = "*.test2"
		queueName   = "test_queue"
		queueName2  = "test_queue2"
		//
		exchange = "test_exchange"
	)

	err = repo.ExchangeDeclare(exchange, queue.ExchangeTypeTopic)
	r.NoError(err)

	err = repo.QueueDeclare(queueName)
	r.NoError(err)
	err = repo.QueueDeclare(queueName2)
	r.NoError(err)

	err = repo.QueueBind(queueName, routingKey, exchange)
	r.NoError(err)
	err = repo.QueueBind(queueName2, routingKey2, exchange)
	r.NoError(err)

	// err = repo.QueueDeclareWithExchange(queueName2, routingKey2, exchange, queue.ExchangeTypeTopic)
	// r.NoError(err)

	jsonData, err := json.Marshal(data)
	r.NoError(err)

	err = repo.Write("test.test2", exchange, jsonData, queue.Json)
	r.NoError(err)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		listenQueue, err := repo.ListenQueue(queueName, "")
		r.NoError(err)

		for m := range listenQueue {
			task := queue.NewQueueData(m)
			task.Ready()

			var dataRow queue.TestData

			err = task.Unmarshal(&dataRow)
			r.NoError(err)

			r.Equal(data.Number, dataRow.Number)
			r.Equal(data.String, dataRow.String)
			break
		}
	}()

	go func() {
		defer wg.Done()

		listenQueue2, err := repo.ListenQueue(queueName2, "")
		r.NoError(err)

		for m := range listenQueue2 {
			task := queue.NewQueueData(m)
			task.Ready()

			var dataRow queue.TestData

			err = task.Unmarshal(&dataRow)
			r.NoError(err)

			r.Equal(data.Number, dataRow.Number)
			r.Equal(data.String, dataRow.String)
			break
		}
	}()

	wg.Wait()
}
