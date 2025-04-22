package models

import (
	"errors"
	"sync"
)

type Balance struct {
	Amount float64
	mu     sync.RWMutex
}

type Account struct {
	UserID  string
	Balance int64
	mu      sync.RWMutex
}

func (b *Balance) Deposit(amount float64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Amount += amount
}

func (b *Balance) Withdraw(amount float64) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.Amount < amount {
		return errors.New("insufficient funds")
	}
	b.Amount -= amount
	return nil
}

func (b *Balance) GetBalance() float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Amount
}

func (a *Account) UpdateBalance(amount int64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
}

func (a *Account) GetBalance() int64 {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.Balance
}
