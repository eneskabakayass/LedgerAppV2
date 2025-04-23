package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"LedgerV2/pkg/services"
)

func CurrentBalanceHandler(w http.ResponseWriter, r *http.Request) {
	balance, err := services.BalanceService.GetCurrentBalance("user-id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance)
}

func HistoricalBalanceHandler(w http.ResponseWriter, r *http.Request) {
	history, err := services.BalanceService.GetBalanceHistory("user-id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

func BalanceAtTimeHandler(w http.ResponseWriter, r *http.Request) {
	at := r.URL.Query().Get("at")
	if at == "" {
		http.Error(w, "missing 'at' query parameter", http.StatusBadRequest)
		return
	}

	timestamp, err := time.Parse(time.RFC3339, at)
	if err != nil {
		http.Error(w, "invalid 'at' format", http.StatusBadRequest)
		return
	}

	balance, err := services.BalanceService.GetBalanceAtTime("user-id", timestamp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(balance)
}
