package internal

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB DB `json:"db" env:"db"`
	QO QO `json:"qo" env:"qo"`
}

type DB struct {
	Table string `json:"table" env:"table"`
}

type QO struct {
	Operation string `json:"operation" env:"operation"`
}

func LoadConfig() Config {
	var config Config

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(fmt.Sprintf("%s/config.json", cwd))
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}

	err = envconfig.Process("nettaton", &config)
	if err != nil {
		panic(err)
	}

	return config
}
