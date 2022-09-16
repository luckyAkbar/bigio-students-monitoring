package config

import (
	"fmt"
	"os"
	"time"
)

var (
	DefaultTokenLength = 35

	DefaultAccessTokenExpiry = time.Hour * 24 * 7
)

func Env() string {
	return os.Getenv("ENV")
}

func LogLevel() string {
	return os.Getenv("LOG_LEVEL")
}

func PostgresDSN() string {
	host := os.Getenv("PG_HOST")
	db := os.Getenv("PG_DATABASE")
	user := os.Getenv("PG_USER")
	pw := os.Getenv("PG_PASSWORD")
	port := os.Getenv("PG_PORT")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pw, db, port)
}

func ServerPort() string {
	cfg := os.Getenv("SERVER_PORT")
	if cfg == "" {
		return "5000"
	}

	return cfg
}
