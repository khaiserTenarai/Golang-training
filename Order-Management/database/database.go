package database

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"order-management/config"
	"order-management/shared/observability"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect establishes a PostgreSQL connection using GORM.
func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Get("DB_HOST", "localhost"),
		config.Get("DB_PORT", "5432"),
		config.Get("DB_USER", "postgres"),
		config.Get("DB_PASSWORD", ""),
		config.Get("DB_NAME", "order_management"),
		config.Get("DB_SSLMODE", "disable"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open PostgreSQL connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get underlying database connection: %w", err)
	}

	// Configure the connection pool. Values can be tuned through environment variables.
	maxOpen := 20
	maxIdle := 5
	if v := config.Get("DB_MAX_OPEN_CONNS", ""); v != "" {
		if n, e := strconv.Atoi(v); e == nil && n > 0 {
			maxOpen = n
		}
	}
	if v := config.Get("DB_MAX_IDLE_CONNS", ""); v != "" {
		if n, e := strconv.Atoi(v); e == nil && n >= 0 {
			maxIdle = n
		}
	}
	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Count database errors centrally across GORM query/write operations.
	_ = db.Callback().Query().After("gorm:query").Register("observability:query_error", func(tx *gorm.DB) {
		if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			observability.IncDatabaseError()
		}
	})
	_ = db.Callback().Create().After("gorm:create").Register("observability:create_error", func(tx *gorm.DB) {
		if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			observability.IncDatabaseError()
		}
	})
	_ = db.Callback().Update().After("gorm:update").Register("observability:update_error", func(tx *gorm.DB) {
		if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			observability.IncDatabaseError()
		}
	})
	_ = db.Callback().Delete().After("gorm:delete").Register("observability:delete_error", func(tx *gorm.DB) {
		if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			observability.IncDatabaseError()
		}
	})

	// Verify that PostgreSQL is reachable.
	if err := sqlDB.Ping(); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("ping PostgreSQL: %w", err)
	}

	return db, nil
}

// Close releases the underlying PostgreSQL connection pool.
func Close(db *gorm.DB) error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
