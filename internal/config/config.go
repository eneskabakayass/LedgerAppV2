package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Port        string `envconfig:"PORT" default:"8080"`
	Environment string `envconfig:"ENV" default:"development"`
	DBURL       string `envconfig:"DATABASE_URL" required:"true"`
}

func Load() *Config {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("env file not loaded: %v", err)
	}

	return &cfg
}
