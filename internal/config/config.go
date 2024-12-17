package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppName     string
	Environment string
	Port        string
	DBDriver    string
	DBAddress   string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		AppName:     getEnv("APP_NAME", "MyTheresa"),
		Environment: getEnv("ENVIRONMENT", "development"),
		Port:        getEnv("PORT", "8080"),

		DBDriver:   getEnv("DB_DRIVER", "postgres"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "admin"),
		DBPassword: getEnv("DB_PASSWORD", "admin"),
		DBName:     getEnv("DB_NAME", "catalog"),
	}

	if cfg.DBDriver == "" || cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBHost == "" || cfg.DBName == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	cfg.DBAddress = fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
		cfg.DBDriver, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBName)
		
	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
