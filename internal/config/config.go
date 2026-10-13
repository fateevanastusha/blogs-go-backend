package config

import (
	"errors"
	"os"
)

type Config struct {
	Port         string
	AuthLogin    string
	AuthPassword string
	LogLevel     string
}

func Load() (Config, error) {
	cfg := Config{
		Port:         getEnv("PORT", "3005"),
		AuthLogin:    os.Getenv("AUTH_LOGIN"),
		AuthPassword: os.Getenv("AUTH_PASSWORD"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
	}
	if cfg.AuthLogin == "" || cfg.AuthPassword == "" {
		return Config{}, errors.New("AUTH_LOGIN and AUTH_PASSWORD must be set")
	}
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
