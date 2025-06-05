package repository

import (
	"errors"
	"time"

	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"
	"gorm.io/gorm"
)

type DayRepository interface {
	Save(day *entities.Day) error
	FindByUserAndDate(userID int, date time.Time) (*entities.Day, error)
}

type dayRepository struct {
	db *gorm.DB
}

func NewDayRepository(db *gorm.DB) DayRepository {
	return &dayRepository{db: db}
}

func (r *dayRepository) Save(day *entities.Day) error {
	return r.db.Create(day).Error
}

func (r *dayRepository) FindByUserAndDate(userID int, date time.Time) (*entities.Day, error) {
	var day entities.Day
	err := r.db.Where("user_id = ? AND date = ?", userID, date).First(&day).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &day, err
}
