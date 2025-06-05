package repository

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entities.User) error
	FindAll() ([]entities.User, error)
	FindByEmail(email string) *entities.User
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindAll() ([]entities.User, error) {
	var users []entities.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindByEmail(email string) *entities.User {
	var user entities.User

	r.db.Where("email = ?", email).Find(&user)

	if user.ID == 0 {
		return nil
	}

	return &user
}
