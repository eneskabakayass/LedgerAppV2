package services

import (
	"LedgerV2/pkg/repositories"
)

type BalanceService struct {
	BalanceRepo repositories.BalanceRepository
	HistoryRepo repositories.BalanceHistoryRepository
}

func NewBalanceService(balanceRepo repositories.BalanceRepository, historyRepo repositories.BalanceHistoryRepository) *BalanceService {
	return &BalanceService{BalanceRepo: balanceRepo, HistoryRepo: historyRepo}
}

func (s *BalanceService) Deposit(userID string, amount float64) error {
	err := s.BalanceRepo.Deposit(userID, amount)
	if err != nil {
		return err
	}
	return s.HistoryRepo.Record(userID, amount, "deposit")
}

func (s *BalanceService) Withdraw(userID string, amount float64) error {
	err := s.BalanceRepo.Withdraw(userID, amount)
	if err != nil {
		return err
	}
	return s.HistoryRepo.Record(userID, amount, "withdraw")
}

func (s *BalanceService) GetBalance(userID string) float64 {
	return s.BalanceRepo.GetBalance(userID)
}
