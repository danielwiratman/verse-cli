package db

import (
	"bible-verse-generator/internal/object"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func New(conf *object.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost, conf.DBPort, conf.DBUser, conf.DBPass, conf.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
