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
