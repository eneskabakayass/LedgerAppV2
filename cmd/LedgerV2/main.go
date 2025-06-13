package main

import (
	"LedgerV2/internal/config"
	"LedgerV2/internal/logger"
	"LedgerV2/internal/server"
	"LedgerV2/pkg/cache"
	"LedgerV2/pkg/services"
	"LedgerV2/pkg/workers"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := config.Load()
	logger.Init(cfg.Environment)
	cache.InitRedis("localhost:6379")

	processor := workers.NewProcessor(5)
	processor.Start()

	txService := services.NewTransactionService()

	server.StartWithService(cfg.Port, txService)

	log.Println("Stopping processor...")
	processor.Stop()
	processor.PrintStats()
}
