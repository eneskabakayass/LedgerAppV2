package repositories

import (
	"LedgerV2/pkg/models"
	"errors"
)

type UserRepository interface {
	Save(user *models.User) error
	FindByID(id string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
}

type InMemoryUserRepository struct {
	users map[string]*models.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*models.User),
	}
}

func (r *InMemoryUserRepository) Save(user *models.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) FindByID(id string) (*models.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepository) FindByUsername(username string) (*models.User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
