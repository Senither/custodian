package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetString(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	return os.Getenv(key)
}
