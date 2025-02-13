package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret string
	AppPort   string
}

func LoadConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		JWTSecret: getEnv("JWT_SECRET", "default_secret"),
		AppPort:   getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
