package repositories

import "LedgerV2/pkg/models"

type TransactionRepository interface {
	Save(transaction *models.Transaction) error
	FindByID(id string) (*models.Transaction, error)
}
