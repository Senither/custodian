package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentConfig struct {
	Application ApplicationConfig
}

type ApplicationConfig struct {
	Name       string
	Descriptor string
	Address    string
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
		},
	}

	return nil
}

func Get() EnvironmentConfig {
	return config
}
