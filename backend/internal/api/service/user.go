package service

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/api/repository"
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain"
	"github.com/VictorzllDev/ZorpPay/backend/internal/valueobjects/email"
)

type UserService interface {
	CreateUser(user *domain.User) error
	GetAllUser() ([]domain.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository: repository}
}

func (s *userService) CreateUser(user *domain.User) error {
	email, err := email.New(user.Email)
	if err != nil {
		return err
	}

	user.Email = email.String()

	return s.repository.Save(user)
}

func (s *userService) GetAllUser() ([]domain.User, error) {
	return s.repository.FindAll()
}
