package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
	Port        string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DatabaseUrl: getEnv("DATABASE_URL", ""),
		Port:        getEnv("PORT", "8080"),
	}
}

func getEnv(key string, defaultValue string) string {
	env := os.Getenv(key)
	if env == "" {
		env = defaultValue
	}
	return env
}
