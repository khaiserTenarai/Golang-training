package database

import (
	"context"
	"fmt"
	"time"

	"employee-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect creates and verifies a PostgreSQL connection using GORM.
func Connect(cfg *config.Config) (*gorm.DB, error) {
	// Step 1: Create the PostgreSQL connection string.
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	// Step 2: Open the database connection using GORM.
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		return nil, fmt.Errorf("open PostgreSQL connection: %w", err)
	}

	// Step 3: Get the underlying database connection pool.
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get SQL connection pool: %w", err)
	}

	// Step 4: Configure connection pool settings.
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	// Step 5: Verify that PostgreSQL is reachable.
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("ping PostgreSQL: %w", err)
	}

	// Step 6: Return the verified GORM database handle.
	return db, nil
}
