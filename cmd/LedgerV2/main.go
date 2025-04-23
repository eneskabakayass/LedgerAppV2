package main

import (
	"LedgerV2/internal/config"
	"LedgerV2/internal/logger"
	"LedgerV2/internal/server"
	"LedgerV2/pkg/services"
	"LedgerV2/pkg/workers"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := config.Load()
	logger.Init(cfg.Environment)

	processor := workers.NewProcessor(5)
	processor.Start()

	txService := services.NewTransactionService()

	server.StartWithService(cfg.Port, txService)

	log.Println("Stopping processor...")
	processor.Stop()
	processor.PrintStats()
}
