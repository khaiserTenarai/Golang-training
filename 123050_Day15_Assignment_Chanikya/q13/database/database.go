package database

import (
	"fmt"

	"ems/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg config.Config) (*gorm.DB, error) {
	// Build the PostgreSQL connection string.
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	// Connect using GORM.
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{TranslateError: true},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
