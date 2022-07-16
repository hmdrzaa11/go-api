package config

import "os"

type Config struct {
	Name        string
	Port        string
	DatabaseURI string
}

func NewConfig() *Config {
	return &Config{
		Name:        env("APP_NAME", "Go Api"),
		Port:        env("APP_PORT", "8000"),
		DatabaseURI: env("DATABASE_URI", ""),
	}
}

func env(key string, defaultValue string) string {
	if value, found := os.LookupEnv(key); found {
		return value
	}
	return defaultValue
}
