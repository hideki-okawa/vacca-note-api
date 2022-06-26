package db

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Okaki030/vacca-note-server/apperr"
	"github.com/Okaki030/vacca-note-server/log"
)

// Connect はdbと接続する関数
func Connect() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	if user == "" {
		return nil, &apperr.ApplicationError{Code: "D001", Err: errors.New("environment variables for DB_USER are not set")}
	}

	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		return nil, &apperr.ApplicationError{Code: "D001", Err: errors.New("environment variables for DB_PASSWORD are not set")}
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		return nil, &apperr.ApplicationError{Code: "D001", Err: errors.New("environment variables for DB_HOST are not set")}
	}

	dbName := os.Getenv("DB_NAME")
	if host == "" {
		return nil, &apperr.ApplicationError{Code: "D001", Err: errors.New("environment variables for DB_NAME are not set")}
	}

	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":3306)/"+dbName+"?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		return nil, &apperr.ApplicationError{Code: "D002", Err: err}
	}

	log.Debugf("connect db")

	return db, nil
}
