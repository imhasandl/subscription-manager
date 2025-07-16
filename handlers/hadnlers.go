package handlers

import (
	"github.com/imhasandl/subscription-manager/database"
)

type Config struct {
	db *database.DB
}

func NewConfig(db *database.DB) *Config {
	return &Config{
		db: db,
	}
}
