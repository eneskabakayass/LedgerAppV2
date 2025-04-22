package repositories

import (
	"errors"
	"sync"
)

type BalanceRepository interface {
	Deposit(userID string, amount float64) error
	Withdraw(userID string, amount float64) error
	GetBalance(userID string) float64
}

type InMemoryBalanceRepository struct {
	data map[string]float64
	mu   sync.RWMutex
}

func NewInMemoryBalanceRepository() *InMemoryBalanceRepository {
	return &InMemoryBalanceRepository{
		data: make(map[string]float64),
	}
}

func (r *InMemoryBalanceRepository) Deposit(userID string, amount float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[userID] += amount
	return nil
}

func (r *InMemoryBalanceRepository) Withdraw(userID string, amount float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.data[userID] < amount {
		return errors.New("insufficient funds")
	}
	r.data[userID] -= amount
	return nil
}

func (r *InMemoryBalanceRepository) GetBalance(userID string) float64 {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.data[userID]
}
