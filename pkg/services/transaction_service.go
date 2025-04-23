package services

import (
	"errors"
	"time"

	"LedgerV2/pkg/models"
)

type TransactionService struct { // <-- T büyük harfle!
	store map[string][]models.Transaction
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		store: make(map[string][]models.Transaction),
	}
}

// Aşağıya metotları ekle (receiver'ı değiştir!):

func (s *TransactionService) Credit(userID string, req models.TransactionRequest) (*models.Transaction, error) {
	tx := models.Transaction{
		ID:        generateID(),
		UserID:    userID,
		Type:      "credit",
		Amount:    req.Amount,
		CreatedAt: time.Now(),
	}
	s.store[userID] = append(s.store[userID], tx)
	return &tx, nil
}

func (s *TransactionService) Debit(userID string, req models.TransactionRequest) (*models.Transaction, error) {
	if req.Amount <= 0 {
		return nil, errors.New("amount must be positive")
	}
	tx := models.Transaction{
		ID:        generateID(),
		UserID:    userID,
		Type:      "debit",
		Amount:    -req.Amount,
		CreatedAt: time.Now(),
	}
	s.store[userID] = append(s.store[userID], tx)
	return &tx, nil
}

func (s *TransactionService) Transfer(fromUserID string, req models.TransferRequest) (*models.Transaction, error) {
	tx := models.Transaction{
		ID:        generateID(),
		UserID:    fromUserID,
		ToUserID:  req.ToUserID,
		Type:      "transfer",
		Amount:    -req.Amount,
		CreatedAt: time.Now(),
	}
	s.store[fromUserID] = append(s.store[fromUserID], tx)
	return &tx, nil
}

func (s *TransactionService) GetHistory(userID string) ([]models.Transaction, error) {
	return s.store[userID], nil
}

func (s *TransactionService) GetByID(userID, id string) (*models.Transaction, error) {
	for _, tx := range s.store[userID] {
		if tx.ID == id {
			return &tx, nil
		}
	}
	return nil, errors.New("transaction not found")
}

func generateID() string {
	return time.Now().Format("20060102150405") + "-" + string(time.Now().UnixNano())
}
