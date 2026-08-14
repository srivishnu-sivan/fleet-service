package config

import (
	"os"
)

type Config struct {
    Port        string
    DatabaseURL string
}

func Load() Config {
	port := os.Getenv("PORT")
	databaseURL := os.Getenv("DATABASE_URL")

	if port == "" {
		port = "8080"
	}

	return Config{
		Port:        port,
		DatabaseURL: databaseURL,
	}
}
