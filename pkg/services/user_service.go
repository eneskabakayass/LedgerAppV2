package services

import (
	"LedgerV2/pkg/cache"
	"LedgerV2/pkg/models"
	"LedgerV2/pkg/repositories"
)

type UserService struct {
	Repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() []*models.User {
	return s.Repo.FindAll()
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	if cachedUser, err := cache.GetUserFromCache(id); err == nil {
		return cachedUser, nil
	}

	user, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	_ = cache.SetUserToCache(user)

	return s.Repo.FindByID(id)
}

func (s *UserService) UpdateUser(id string, data models.UserUpdateRequest) (*models.User, error) {
	user, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Email = data.Email
	user.Role = data.Role

	err = s.Repo.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(id string) error {
	return s.Repo.Delete(id)
}
