package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB returns *grom.DB driver
func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseDB(db *gorm.DB) error {
	postgres, _ := db.DB()
	err := postgres.Close()
	if err != nil {
		return err
	}
	return nil
}
