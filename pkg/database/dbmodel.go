package database

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DBModel struct {
	Conn *sql.DB
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
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name: os.Getenv("DB_NAME"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
	}

	if dbModel.User == "" || dbModel.Password == "" || dbModel.Name == "" ||
		dbModel.Host == "" || dbModel.Port == "" {
		return nil, errors.New("DBModel Error")
	}

	dsn := dbModel.User + ":" + dbModel.Password + "@tcp(" + dbModel.Host + ":" + dbModel.Port + ")/" + dbModel.Name + "?parseTime=true"

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	dbModel.Conn = conn

	return dbModel, nil
}

func (db *DBModel) GetConn() *sql.DB {
	return db.Conn
}
