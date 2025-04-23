package handlers

import (
	"encoding/json"
	"net/http"

	"LedgerV2/pkg/models"
	"LedgerV2/pkg/services"
	"github.com/go-chi/chi/v5"
)

type TransactionHandler struct {
	Service *services.TransactionService
}

func (h *TransactionHandler) CreditHandler(w http.ResponseWriter, r *http.Request) {
	var tx models.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	transaction, err := h.Service.Credit("user-id", tx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) DebitHandler(w http.ResponseWriter, r *http.Request) {
	var tx models.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	transaction, err := h.Service.Debit("user-id", tx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) TransferHandler(w http.ResponseWriter, r *http.Request) {
	var tx models.TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	transaction, err := h.Service.Transfer("user-id", tx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) TransactionHistoryHandler(w http.ResponseWriter, r *http.Request) {
	history, err := h.Service.GetHistory("user-id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(history)
}

func (h *TransactionHandler) GetTransactionByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	tx, err := h.Service.GetByID("user-id", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(tx)
}
