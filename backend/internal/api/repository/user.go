package repository

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *domain.User) error
	FindAll() ([]domain.User, error)
	FindByEmail(email string) *domain.User
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindByEmail(email string) *domain.User {
	var user domain.User

	r.db.Where("email = ?", email).Find(&user)

	if user.ID == 0 {
		return nil
	}

	return &user
}
