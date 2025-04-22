package services

import (
	"errors"

	"LedgerV2/pkg/models"
	"LedgerV2/pkg/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (s *UserService) Register(user *models.User, plainPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.UserRepo.Save(user)
}

func (s *UserService) Authenticate(username, password string) (*models.User, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (s *UserService) Authorize(user *models.User, requiredRole string) bool {
	return user.Role == requiredRole
}
