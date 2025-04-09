package models

import (
	"sync"
)

type Transaction struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
	mu     sync.Mutex
	UserID int
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
