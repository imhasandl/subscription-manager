package database

import (
	"database/sql"
	"errors"
	"log"
)

type DB struct {
	*sql.DB
}

func InitDatabase(db_url string) (*DB, error) {
	db, err := sql.Open("postgres", db_url)
	if err != nil {
		return nil, errors.New("failed to open database")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.New("failed to ping database")
	}

	log.Println("Successfully connected to database")
	return &DB{db}, nil
}

func (db *DB) Close() error {
	return db.DB.Close()
}
