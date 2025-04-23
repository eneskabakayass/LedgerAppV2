package services

import (
	"LedgerV2/pkg/models"
	"time"
)

type balanceService struct{}

var BalanceService = balanceService{}

func (s balanceService) GetCurrentBalance(userID string) (*models.Balance, error) {
	return &models.Balance{
		UserID: userID,
		Amount: 1250.50,
		Date:   time.Now(),
	}, nil
}

func (s balanceService) GetBalanceHistory(userID string) ([]models.Balance, error) {
	return []models.Balance{
		{UserID: userID, Amount: 1000.00, Date: time.Now().AddDate(0, -1, 0)},
		{UserID: userID, Amount: 1200.00, Date: time.Now().AddDate(0, -2, 0)},
		{UserID: userID, Amount: 1250.50, Date: time.Now()},
	}, nil
}

func (s balanceService) GetBalanceAtTime(userID string, t time.Time) (*models.Balance, error) {
	return &models.Balance{
		UserID: userID,
		Amount: 1111.11,
		Date:   t,
	}, nil
}
