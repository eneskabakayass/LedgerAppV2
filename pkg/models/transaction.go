package models

import "time"

type Transaction struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Type      string    `json:"type"` // credit, debit, transfer
	Amount    float64   `json:"amount"`
	ToUserID  string    `json:"to_user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type TransactionRequest struct {
	Amount float64 `json:"amount"`
}

type TransferRequest struct {
	ToUserID string  `json:"to_user_id"`
	Amount   float64 `json:"amount"`
}
