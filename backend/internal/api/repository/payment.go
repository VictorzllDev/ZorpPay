package repository

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Save(payment *entities.Payment) error
	CountUserDays(userId int) (int, error)
	GetUserTotalPayments(userId int) (int, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Save(payment *entities.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) CountUserDays(userId int) (int, error) {
	var count int64
	result := r.db.Model(&entities.Day{}).
		Where("user_id = ?", userId).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}

func (r *paymentRepository) GetUserTotalPayments(userId int) (int, error) {
	var totalUsed int
	err := r.db.Model(&entities.Payment{}).
		Select("COALESCE(SUM(amount), 0)").
		Where("user_id = ?", userId).
		Scan(&totalUsed).Error

	return totalUsed, err
}
