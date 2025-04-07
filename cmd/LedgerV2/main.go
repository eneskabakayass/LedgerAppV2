package main

import (
	"LedgerV2/internal/config"
	"LedgerV2/internal/logger"
	"LedgerV2/internal/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	cfg := config.Load()
	logger.Init(cfg.Environment)
	server.Start(cfg.Port)
}
