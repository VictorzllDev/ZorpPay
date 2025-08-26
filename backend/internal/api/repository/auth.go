package repository

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Save(user *entities.User) error
	FindByEmail(email string) *entities.User
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Save(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) FindByEmail(email string) *entities.User {
	user := &entities.User{}

	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil
	}

	return user
}
