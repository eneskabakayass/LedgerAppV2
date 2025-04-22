package server

import (
	"LedgerV2/pkg/models"
	"LedgerV2/pkg/services"
	"context"
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func StartWithService(port string, txService *services.TransactionService) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello Ledger Application"))
		if err != nil {
			log.Error().Err(err).Msg("Failed to write response")
		}
	})

	mux.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var tx models.Transaction
		if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
			log.Error().Err(err).Msg("Invalid transaction payload")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if tx.ID == "" {
			tx.ID = strconv.Itoa(int(time.Now().UnixNano()))
		}

		log.Info().Str("user", tx.FromUserID).Int64("amount", int64(tx.Amount)).Msg("Transaction received")

		txService.SubmitTransaction(&tx)

		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write([]byte("Transaction received"))
	})

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	go func() {
		log.Info().Msgf("Server started on port %s", port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Server error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("The server did not shut down properly")
	} else {
		log.Info().Msg("The server was shut down properly")
	}
}
