package services

import (
	"LedgerV2/pkg/models"
	"LedgerV2/pkg/repositories"
	"errors"
)

type TransactionService struct {
	BalanceRepo repositories.BalanceRepository
}

func NewTransactionService(balanceRepo repositories.BalanceRepository) *TransactionService {
	return &TransactionService{BalanceRepo: balanceRepo}
}

func (s *TransactionService) Credit(userID string, amount float64) error {
	return s.BalanceRepo.Deposit(userID, amount)
}

func (s *TransactionService) Debit(userID string, amount float64) error {
	return s.BalanceRepo.Withdraw(userID, amount)
}

func (s *TransactionService) Transfer(fromID, toID string, amount float64) error {
	if err := s.BalanceRepo.Withdraw(fromID, amount); err != nil {
		return err
	}
	if err := s.BalanceRepo.Deposit(toID, amount); err != nil {
		_ = s.BalanceRepo.Deposit(fromID, amount)
		return errors.New("transfer failed, rollback applied")
	}
	return nil
}

func (s *TransactionService) SubmitTransaction(tx *models.Transaction) error {
	switch tx.Type {
	case "credit":
		return s.Credit(tx.ToUserID, tx.Amount)
	case "debit":
		return s.Debit(tx.FromUserID, tx.Amount)
	case "transfer":
		return s.Transfer(tx.FromUserID, tx.ToUserID, tx.Amount)
	default:
		return errors.New("invalid transaction type")
	}
}
