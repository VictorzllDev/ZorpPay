package service

import (
	"errors"
	"fmt"

	"github.com/VictorzllDev/ZorpPay/backend/internal/api/repository"
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/valueobjects/email"
	"github.com/VictorzllDev/ZorpPay/backend/internal/pkg/security"
)

type UserService interface {
	CreateUser(user *entities.User) error
	GetAllUser() ([]entities.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository: repository}
}

func (s *userService) CreateUser(user *entities.User) error {
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

func (s *userService) GetAllUser() ([]entities.User, error) {
	return s.repository.FindAll()
}
