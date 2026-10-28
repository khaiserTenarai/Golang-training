package database

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"time"

	"employee-api/config"
	"employee-api/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect opens PostgreSQL and checks the connection.
func Connect(
	cfg *config.Config,
	secrets *config.Secrets,
) (*gorm.DB, error) {
	// Encode the password safely, including special characters.
	dsnURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.DBUser, secrets.DBPassword),
		Host:   net.JoinHostPort(cfg.DBHost, cfg.DBPort),
		Path:   cfg.DBName,
	}

	// Add the PostgreSQL SSL mode.
	query := url.Values{}
	query.Set("sslmode", cfg.DBSSLMode)
	dsnURL.RawQuery = query.Encode()

	// Open the database through GORM.
	db, err := gorm.Open(
		postgres.Open(dsnURL.String()),
		&gorm.Config{},
	)
	if err != nil {
		return nil, fmt.Errorf("open PostgreSQL: %w", err)
	}

	// Get the underlying SQL connection pool.
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get SQL connection pool: %w", err)
	}

	// Set connection pool limits.
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	// Verify database connectivity with a five-second timeout.
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

// Migrate creates the employee table if it does not exist.
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.Employee{}); err != nil {
		return fmt.Errorf("migrate employee table: %w", err)
	}

	return nil
}