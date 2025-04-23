package services

import (
	"LedgerV2/pkg/models"
	"LedgerV2/pkg/repositories"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var AuthService = authService{}

type authService struct{}

var jwtSecret = []byte("super-secret-key")

func (s authService) Login(req models.LoginRequest) (*models.TokenResponse, error) {
	user, err := repositories.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if user.Password != req.Password+"_hashed" {
		return nil, errors.New("invalid credentials")
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.TokenResponse{
		AccessToken:  token,
		RefreshToken: "dummy-refresh-token",
	}, nil
}

func (s authService) Refresh() (*models.TokenResponse, error) {
	token, err := generateJWT("some-user-id")
	if err != nil {
		return nil, err
	}
	return &models.TokenResponse{
		AccessToken:  token,
		RefreshToken: "new-dummy-refresh-token",
	}, nil
}

func generateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (s authService) Register(req models.RegisterRequest) (*models.User, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	exists, _ := repositories.UserRepo.ExistsByEmail(req.Email)
	if exists {
		return nil, errors.New("user already exists")
	}

	user := &models.User{
		ID:       uuid.NewString(),
		Email:    req.Email,
		Password: req.Password + "_hashed",
		Role:     "user",
	}

	if err := repositories.UserRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
