package entities

import (
	"time"
)

type Payment struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Amount      float64   `gorm:"not null;default:0" json:"amount"` // Valor pago (ex: 140.00)
	PaymentDate time.Time `gorm:"not null" json:"payment_date"`     // Data do pagamento
	UserID      int       `gorm:"not null;index" json:"user_id"`    // Chave estrangeira
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
