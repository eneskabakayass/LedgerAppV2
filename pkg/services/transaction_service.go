package services

import (
	"LedgerV2/pkg/models"
	"LedgerV2/pkg/workers"
)

type TransactionService struct {
	Processor *workers.Processor
}

func NewTransactionService(p *workers.Processor) *TransactionService {
	return &TransactionService{Processor: p}
}

func (s *TransactionService) SubmitTransaction(tx *models.Transaction) {
	s.Processor.AddTransaction(tx)
}
