package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Port        string `env:"PORT" envDefault:"3333"`
	Env         string `env:"ENV" envDefault:"development"`
	DatabaseURL string `env:"DATABASE_URL,required"`
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("[WARN] Error loading .env file: %v", err)
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("config parsing failed: %w", err)
	}

	return &cfg, nil
}
