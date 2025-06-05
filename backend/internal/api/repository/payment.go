package repository

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Save(payment *entities.Payment) error
	CountDays() (int, error)
	GetTotalUsedInPayments() (int, error)
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

func (r *paymentRepository) CountDays() (int, error) {
	var count int64
	result := r.db.Model(&entities.Day{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}

func (r *paymentRepository) GetTotalUsedInPayments() (int, error) {
	var totalUsed int
	err := r.db.Model(&entities.Payment{}).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&totalUsed).Error

	return totalUsed, err
}
