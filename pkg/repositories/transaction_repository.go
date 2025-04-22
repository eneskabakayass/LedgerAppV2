package repositories

import (
	"LedgerV2/pkg/models"
	"errors"
)

type TransactionRepository interface {
	Save(transaction *models.Transaction) error
	FindByID(id string) (*models.Transaction, error)
}

type InMemoryTransactionRepository struct {
	data map[string]*models.Transaction
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{
		data: make(map[string]*models.Transaction),
	}
}

func (r *InMemoryTransactionRepository) Save(tx *models.Transaction) error {
	r.data[tx.ID] = tx
	return nil
}

func (r *InMemoryTransactionRepository) FindByID(id string) (*models.Transaction, error) {
	tx, ok := r.data[id]
	if !ok {
		return nil, errors.New("transaction not found")
	}
	return tx, nil
}
