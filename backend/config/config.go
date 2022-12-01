package config

import (
	"fmt"
	"os"

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
}

// NewConfig init and return project config
func NewConfig(confPath string) (Config, error) {
	var c = Config{}
	err := configor.Load(&c, confPath)
	return c, err
}

// PostgresURL connect to callcenter's postgres db
func (c Config) PostgresURL() string {
	if os.Getenv("DB_URL") != "" {
		c.Postgres.URL = os.Getenv("DB_URL")
	}
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Postgres.User, c.Postgres.Password, c.Postgres.URL, c.Postgres.User)
}

func (c Config) ApiURL() string {
	return fmt.Sprintf("%s:%d", c.API.IP, c.API.Port)
}
