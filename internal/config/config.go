package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/joho/godotenv"
)

type (
	// Config - cтруктура для хранения конфигурации приложения.
	Config struct {
		// Параметры бд
		Postgres struct {
			User 		string `envconfig:"PG_USER" required:"true"`
			Password 	string `envconfig:"PG_PASSWORD" required:"true"`
			Database 	string `envconfig:"PG_DB" required:"true"`
			HostPort 	string `envconfig:"PG_HOST_PORT" required:"true"`
		}

		// Параметры для HTTP
		HTTP struct {
			Port string `envconfig:"HTTP_PORT" required:"true"`
		}

		// Параметры для JWT
		JWT struct {
			Secret string `envconfig:"JWT_SECRET" required:"true"`
		}
	}
)

// GetConfigFromEnv - загружает конфигурации из .env файла и переменных окружения.
func GetConfigFromEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load the .env file: %s\n", err.Error())
	}

	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
