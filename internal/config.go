package internal

import (
	"os"
)

// Config ...
type Config struct {
	DB DB `json:"db" env:"db"`
}

// DB ...
type DB struct {
	Table string `json:"table" env:"table"`
}

// LoadConfig ...
func LoadConfig() Config {
	var config Config

	// cwd, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }

	// file, err := os.Open(fmt.Sprintf("%s/config.json", cwd))
	// if err != nil {
	// 	panic(err)
	// }

	// err = json.NewDecoder(file).Decode(&config)
	// if err != nil {
	// 	panic(err)
	// }

	// err = envconfig.Process("nettaton", &config)
	// if err != nil {
	// 	panic(err)
	// }

	config.DB.Table = os.Getenv("NETTATON_TABLE")

	return config
}
