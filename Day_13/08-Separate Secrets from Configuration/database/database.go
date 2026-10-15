package database

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"time"

	"employee-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect opens and verifies a PostgreSQL connection.
func Connect(
	cfg *config.Config,
	secrets *config.Secrets,
) (*gorm.DB, error) {
	// Step 1: Build a PostgreSQL URL safely.
	dsnURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.DBUser, secrets.DBPassword),
		Host:   net.JoinHostPort(cfg.DBHost, cfg.DBPort),
		Path:   cfg.DBName,
	}

	query := url.Values{}
	query.Set("sslmode", cfg.DBSSLMode)
	dsnURL.RawQuery = query.Encode()

	// Step 2: Open the database using GORM.
	db, err := gorm.Open(
		postgres.Open(dsnURL.String()),
		&gorm.Config{},
	)
	if err != nil {
		return nil, fmt.Errorf("open PostgreSQL connection: %w", err)
	}

	// Step 3: Get the underlying SQL connection pool.
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get SQL connection pool: %w", err)
	}

	// Step 4: Configure connection pool limits.
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	// Step 5: Verify connectivity with a timeout.
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("ping PostgreSQL: %w", err)
	}

	return db, nil
}
