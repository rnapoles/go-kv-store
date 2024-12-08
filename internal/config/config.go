package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBType   string
	DBPath   string
	HTTPPort string
}

func LoadConfig() (*Config, error) {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Read environment variables
	return &Config{
		DBType:   os.Getenv("DB_TYPE"),
		DBPath:   os.Getenv("DB_PATH"),
		HTTPPort: os.Getenv("HTTP_PORT"),
	}, nil
}
