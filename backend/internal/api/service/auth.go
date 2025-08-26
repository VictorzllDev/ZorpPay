package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/VictorzllDev/ZorpPay/backend/internal/api/repository"
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/valueobjects/email"
	"github.com/VictorzllDev/ZorpPay/backend/internal/pkg/security"
)

type AuthService interface {
	SignUp(req *entities.User) error
	SignIn(req *entities.User) (*entities.Auth, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(repository repository.AuthRepository) AuthService {
	return &authService{repository: repository}
}

func (s *authService) SignUp(user *entities.User) error {
	email, err := email.New(user.Email)
	if err != nil {
		return err
	}

	existingUser := s.repository.FindByEmail(email.String())
	if existingUser != nil {
		return errors.New("user already exists")
	}

	hash, err := security.GenerateHash(user.Password)
	if err != nil {
		return fmt.Errorf("password hashing failed: %w", err)
	}
	user.Password = string(hash)

	return s.repository.Save(user)
}

func (s *authService) SignIn(user *entities.User) (*entities.Auth, error) {
	existingUser := s.repository.FindByEmail(user.Email)
	if existingUser == nil {
		return nil, errors.New("user not found")
	}

	if !security.CompareHash(existingUser.Password, user.Password) {
		return nil, errors.New("invalid password")
	}

	token, err := security.NewJWT().GenerateToken(strconv.Itoa(existingUser.ID))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &entities.Auth{
		Token: token,
	}, nil
}
