package repositories

import "LedgerV2/pkg/models"

type UserRepository interface {
	Save(user *models.User) error
	FindByID(id string) (*models.User, error)
}
