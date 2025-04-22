package models

import (
	"sync"
)

type Transaction struct {
	ID         string  `json:"id"`
	FromUserID string  `json:"from_user_id,omitempty"`
	ToUserID   string  `json:"to_user_id,omitempty"`
	Amount     float64 `json:"amount"`
	Type       string  `json:"type"`
	Status     string  `json:"status"`
	mu         sync.Mutex
}

func (t *Transaction) SetStatus(status string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Status = status
}

func (t *Transaction) IsCompleted() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.Status == "completed"
}
