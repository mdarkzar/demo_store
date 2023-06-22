package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/configor"
)

// Config конфиг
type Config struct {
	Postgres struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		URL      string `yaml:"url"`
		Port     string `yaml:"port"`
	} `yaml:"postgres"`
	API struct {
		IP   string `yaml:"ip"`
		Port int    `yaml:"port"`
	} `yaml:"api"`
	Queue struct {
		ReconnectSeconds int `yaml:"reconnect_seconds"`
	} `yaml:"queue"`
}

// NewConfig инициализация конфига, считывания его с файла
func NewConfig(confPath string) (Config, error) {
	var c = Config{}
	err := configor.Load(&c, confPath)
	return c, err
}

// PostgresURL сборка url для подключение к pg
func (c Config) PostgresURL() string {
	if os.Getenv("DB_URL") != "" {
		c.Postgres.URL = os.Getenv("DB_URL")
	}
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Postgres.User, c.Postgres.Password, c.Postgres.URL, c.Postgres.User)
}

func (c Config) ApiURL() string {
	return fmt.Sprintf("%s:%d", c.API.IP, c.API.Port)
}

// RabbitMQConnectURL подключение к rabbitmq
func (c Config) RabbitMQConnectURL() string {
	if os.Getenv("RABBIT_URL") != "" {
		return fmt.Sprintf("amqp://%s/", os.Getenv("RABBIT_URL"))
	}

	log.Fatalln("отсутствует RABBIT_URL")
	return ""
}

func (c Config) ReconnectQueue() time.Duration {
	return time.Duration(c.Queue.ReconnectSeconds) * time.Second
}
