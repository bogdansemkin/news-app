package config

import (
	"github.com/joho/godotenv"
	"github.com/vrischmann/envconfig"
	"log"
)

type Config struct {
	HTTP struct {
		Host string `envconfig:"HOST"`
		Port string `envconfig:"PORT"`
	}
	DB struct {
		PostgresConn string `envconfig:"PostgresDBconn"`
	}
}

func LoadConfig() (*Config, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	var cfg Config
	if err := envconfig.Init(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
