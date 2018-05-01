package internal

import (
	"os"
)

// Config stores all the config information for the app.
type Config struct {
	DB DB `json:"db" env:"db"`
}

// DB has all the config info for the database.
type DB struct {
	Table string `json:"table" env:"table"`
}

// LoadConfig loads the config into the config object.
func LoadConfig() Config {
	var config Config

	table := os.Getenv("NETTATON_TABLE")
	if len(table) == 0 {
		panic("NETTATON_TABLE not set.")
	}
	config.DB.Table = table

	return config
}
