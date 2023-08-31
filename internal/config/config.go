package config

import (
	"avito-backend-trainee-assignment-2023/pkg/logging"
	"fmt"
	"os"
	"strconv"
)

var logger = logging.GetLogger()

func GetPostgresDSN() string {
	user, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		logger.Fatal("postgres user is not set in env!")
	}
	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		logger.Fatal("postgres password is not set in env!")
	}
	host, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		logger.Fatal("postgres host is not set in env!")
	}
	portStr, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		logger.Fatal("postgres port is not set in env!")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		logger.Fatal("postgres port is incorrect: ", err)
	}
	db, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		logger.Fatal("postgres db is not set in env!")
	}
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		db,
	)
}

func GetAPIAddress() string {
	host, ok := os.LookupEnv("API_HOST")
	if !ok {
		logger.Fatal("api host is not set in env!")
	}
	portStr, ok := os.LookupEnv("API_PORT")
	if !ok {
		logger.Fatal("api port is not set in env!")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		logger.Fatal("api port is incorrect: ", err)
	}
	return fmt.Sprintf("%s:%d", host, port)
}
