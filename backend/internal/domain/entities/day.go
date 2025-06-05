package entities

import (
	"time"
)

type Day struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Date      time.Time `gorm:"not null;uniqueIndex:idx_user_date" json:"date"` // Data do dia letivo
	UserID    int       `gorm:"not null;index" json:"user_id"`                  // Chave estrangeira
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
