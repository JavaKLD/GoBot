package database

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
)

type DBModel struct {
	User string
	Password string
	Name string
	Host string
	Port string
}

func LoadDBModel() (*DBModel, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbModel := &DBModel{
		User:     "DB_USER",
		Password: "DB_PASSWORD",
		Name: "DB_NAME",
		Host:     "DB_HOST",
		Port:     "DB_PORT",
	}

	if dbModel.User == "" || dbModel.Password == "" || dbModel.Name == "" ||
		dbModel.Host == "" || dbModel.Port == "" {
		return nil, errors.New("DBModel Error")
	}
	return dbModel, nil
}
