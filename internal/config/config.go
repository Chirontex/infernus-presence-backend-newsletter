package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	ClientToken   string
	Database      DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load() (*Config, error) {
	// Load .env file if exists (for local development)
	_ = godotenv.Load()

	cfg := &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		ClientToken:   getEnv("CLIENT_TOKEN", ""),
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			Database: getEnv("DB_NAME", "newsletter"),
		},
	}

	if cfg.ClientToken == "" {
		return nil, fmt.Errorf("CLIENT_TOKEN is required")
	}

	return cfg, nil
}

func (d DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		d.User, d.Password, d.Host, d.Port, d.Database)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
