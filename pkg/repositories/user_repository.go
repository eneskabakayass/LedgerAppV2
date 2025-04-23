package repositories

import (
	"LedgerV2/pkg/models"
	"github.com/pkg/errors"
)

type UserRepository interface {
	ExistsByEmail(email string) (bool, error)
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	Save(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	FindAll() []*models.User
	FindByID(id string) (*models.User, error)
	Delete(id string) error
}

type userRepository struct {
	store map[string]*models.User
}

var UserRepo = userRepository{
	store: make(map[string]*models.User),
}

func (r userRepository) ExistsByEmail(email string) (bool, error) {
	for _, u := range r.store {
		if u.Email == email {
			return true, nil
		}
	}
	return false, nil
}

func (r userRepository) Create(user *models.User) error {
	r.store[user.ID] = user
	return nil
}

func (r userRepository) FindByEmail(email string) (*models.User, error) {
	for _, u := range r.store {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r userRepository) Save(user *models.User) error {
	r.store[user.ID] = user
	return nil
}

func (r userRepository) FindByUsername(username string) (*models.User, error) {
	return r.FindByEmail(username)
}

func (r userRepository) FindAll() []*models.User {
	var users []*models.User
	for _, u := range r.store {
		users = append(users, u)
	}
	return users
}

func (r userRepository) FindByID(id string) (*models.User, error) {
	user, ok := r.store[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r userRepository) Delete(id string) error {
	if _, ok := r.store[id]; !ok {
		return errors.New("user not found")
	}
	delete(r.store, id)
	return nil
}
