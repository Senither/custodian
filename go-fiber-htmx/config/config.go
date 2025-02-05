package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	Application ApplicationConfig
	Database    DatabaseConfig
}

type ApplicationConfig struct {
	Name       string
	Descriptor string
	Address    string
	Debug      bool
}

type DatabaseConfig struct {
	Url string
}

var config EnvironmentConfig

func LoadConfig() error {
	err := godotenv.Load(".env")

	if err != nil {
		return err
	}

	config = EnvironmentConfig{
		Application: ApplicationConfig{
			Name:       os.Getenv("APP_NAME"),
			Descriptor: os.Getenv("APP_DESCRIPTOR"),
			Address:    os.Getenv("APP_ADDR"),
			Debug:      os.Getenv("APP_DEBUG") == "true",
		},
		Database: DatabaseConfig{
			Url: os.Getenv("DATABASE_URL"),
		},
	}

	return nil
}

func Get() EnvironmentConfig {
	return config
}
