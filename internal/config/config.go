package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Port        string `envconfig:"PORT" default:"3306"`
	Environment string `envconfig:"ENV" default:"development"`
	DBURL       string `envconfig:"DATABASE_URL" required:"root:12345@tcp(db:3306)/ledger_app_db_v2"`
}

func Load() *Config {
	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("env file not loaded: %v", err)
	}

	return &cfg
}
