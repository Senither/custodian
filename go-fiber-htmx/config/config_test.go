package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createDummyEnvFile(t *testing.T) func() {
	envContent := `
	APP_NAME=TestApp
	APP_DESCRIPTOR=Test Descriptor
	APP_DEBUG=true
	APP_ADDR=:8080
	DATABASE_URL=test_database.sqlite
	`

	err := os.WriteFile(".env.test", []byte(envContent), 0644)
	assert.Nil(t, err)

	return func() {
		os.Remove(".env.test")
	}
}

func TestLoadConfig(t *testing.T) {
	t.Run("it loads configuration from .env file", func(t *testing.T) {
		cleanup := createDummyEnvFile(t)
		defer cleanup()

		err := LoadConfig(".env.test")
		assert.Nil(t, err)

		expectedConfig := EnvironmentConfig{
			Application: ApplicationConfig{
				Name:       "TestApp",
				Descriptor: "Test Descriptor",
				Address:    ":8080",
				Debug:      true,
			},
			Database: DatabaseConfig{
				Url: "test_database.sqlite",
			},
		}

		assert.Equal(t, expectedConfig, Get())
	})

	t.Run("it returns an error if .env file is missing", func(t *testing.T) {
		err := LoadConfig(".env.invalid")
		assert.NotNil(t, err)
	})
}

func TestGet(t *testing.T) {
	t.Run("it returns the loaded configuration", func(t *testing.T) {
		cleanup := createDummyEnvFile(t)
		defer cleanup()

		expectedConfig := EnvironmentConfig{
			Application: ApplicationConfig{
				Name:       "TestApp",
				Descriptor: "Test Descriptor",
				Address:    ":8080",
				Debug:      true,
			},
			Database: DatabaseConfig{
				Url: "test_database.sqlite",
			},
		}

		err := LoadConfig(".env.test")
		assert.Nil(t, err)
		assert.Equal(t, expectedConfig, Get())
	})
}
