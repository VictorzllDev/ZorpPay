package database

import (
	"github.com/VictorzllDev/ZorpPay/backend/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Executar migrations
	if err := entities.Migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}
