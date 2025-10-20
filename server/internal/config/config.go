package config

import (
	"os"
)

var (
	POSTGRES_URL = getEnv("HALO_DB_DSN")
)

func getEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return "None"
}

func GetPostgresUrl() string {
	return os.Getenv("HALO_DB_DSN")
}
