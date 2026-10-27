package database

import (
	"day19/employee-service/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Open creates a GORM connection to PostgreSQL.
func Open(cfg config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connect to PostgreSQL: %w", err)
	}
	return db, nil
}
