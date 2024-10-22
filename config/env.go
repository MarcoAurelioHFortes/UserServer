package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBDatabase string
	DBUsername string
	DBPassword string
	DBSchema   string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		DBHost:     getEnv("DB_HOST", "http://localhost"),
		DBPort:     getEnv("DB_PORT", "8080"),
		DBDatabase: getEnv("DB_DATABASE", "master"),
		DBUsername: getEnv("DB_USERNAME", "admin"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBSchema:   getEnv("DB_SCHEMA", "public"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
