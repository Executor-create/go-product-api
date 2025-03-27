package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return nil
}

func GetEnvVariable(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("Warning: environment variable %s not set", key)
	}

	return value
}
