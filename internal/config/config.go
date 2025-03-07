package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PGHost     string
	PGPort     string
	PGUser     string
	PGPassword string
	PGName     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

	return &Config{
		PGHost:     os.Getenv("PG_HOST"),
		PGPort:     os.Getenv("PG_PORT"),
		PGUser:     os.Getenv("PG_USER"),
		PGPassword: os.Getenv("PG_PASSWORD"),
		PGName:     os.Getenv("PG_NAME"),
	}
}
