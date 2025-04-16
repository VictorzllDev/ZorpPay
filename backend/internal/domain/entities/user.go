package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
