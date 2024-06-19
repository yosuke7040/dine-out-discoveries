package database

import (
	"os"
	"time"
)

type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string

	ctxTimeout time.Duration
}

func newConfigMySql() *config {
	return &config{
		host:     os.Getenv("MYSQL_HOST"),
		database: os.Getenv("MYSQL_DATABASE"),
		port:     os.Getenv("MYSQL_PORT"),
		// driver:     os.Getenv("MYSQL_DRIVER"),
		user:       os.Getenv("MYSQL_USER"),
		password:   os.Getenv("MYSQL_PASSWORD"),
		ctxTimeout: 10 * time.Second,
	}
}

func newConfigPostgres() *config {
	return &config{
		host:     os.Getenv("POSTGRES_HOST"),
		database: os.Getenv("POSTGRES_DATABASE"),
		port:     os.Getenv("POSTGRES_PORT"),
		driver:   os.Getenv("POSTGRES_DRIVER"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}
}
