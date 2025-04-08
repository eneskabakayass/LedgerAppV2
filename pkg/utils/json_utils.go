package utils

import (
	"LedgerV2/pkg/models"
	"encoding/json"
)

func MarshalUser(user *models.User) ([]byte, error) {
	return json.Marshal(user)
}

func UnmarshalUser(data []byte) (*models.User, error) {
	var user models.User
	err := json.Unmarshal(data, &user)
	return &user, err
}

func MarshalTransaction(transaction *models.Transaction) ([]byte, error) {
	return json.Marshal(transaction)
}

func UnmarshalTransaction(data []byte) (*models.Transaction, error) {
	var transaction models.Transaction
	err := json.Unmarshal(data, &transaction)
	return &transaction, err
}

func MarshalBalance(balance *models.Balance) ([]byte, error) {
	return json.Marshal(balance)
}

func UnmarshalBalance(data []byte) (*models.Balance, error) {
	var balance models.Balance
	err := json.Unmarshal(data, &balance)
	return &balance, err
}
