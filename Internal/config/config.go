package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	StaticDir string
	Port      string
}

func Load(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return &Config{}, err
	}
	cfg := Config{
		StaticDir: os.Getenv("STATIC_DIR"),
		Port:      os.Getenv("API_PORT"),
	}
	return &cfg, nil
}
