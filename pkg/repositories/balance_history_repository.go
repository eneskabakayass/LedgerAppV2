package repositories

import (
	"fmt"
	"time"
)

type BalanceHistoryEntry struct {
	UserID string
	Amount float64
	Action string
	Time   time.Time
}

type BalanceHistoryRepository interface {
	Record(userID string, amount float64, action string) error
	GetHistory(userID string) []BalanceHistoryEntry
}

type InMemoryBalanceHistoryRepository struct {
	history map[string][]BalanceHistoryEntry
}

func NewInMemoryBalanceHistoryRepository() *InMemoryBalanceHistoryRepository {
	return &InMemoryBalanceHistoryRepository{
		history: make(map[string][]BalanceHistoryEntry),
	}
}

func (r *InMemoryBalanceHistoryRepository) Record(userID string, amount float64, action string) error {
	entry := BalanceHistoryEntry{
		UserID: userID,
		Amount: amount,
		Action: action,
		Time:   time.Now(),
	}
	r.history[userID] = append(r.history[userID], entry)
	fmt.Printf("History: %s - %.2f (%s)\n", userID, amount, action)
	return nil
}

func (r *InMemoryBalanceHistoryRepository) GetHistory(userID string) []BalanceHistoryEntry {
	return r.history[userID]
}
