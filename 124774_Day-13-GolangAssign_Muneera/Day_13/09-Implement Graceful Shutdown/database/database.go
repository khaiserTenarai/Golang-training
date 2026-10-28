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
	// Build the database URL.
	// url.UserPassword safely encodes special characters in the password.
	dsnURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.DBUser, secrets.DBPassword),
		Host:   net.JoinHostPort(cfg.DBHost, cfg.DBPort),
		Path:   cfg.DBName,
	}

	// Add the PostgreSQL SSL setting to the URL.
	query := url.Values{}
	query.Set("sslmode", cfg.DBSSLMode)
	dsnURL.RawQuery = query.Encode()

	// Open the database through GORM.
	db, err := gorm.Open(
		postgres.Open(dsnURL.String()),
		&gorm.Config{},
	)
	if err != nil {
		return nil, fmt.Errorf("open PostgreSQL connection: %w", err)
	}

	// Get the underlying database connection pool.
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get SQL connection pool: %w", err)
	}

	// Limit the number of open and idle database connections.
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)

	// Reuse connections for a limited time.
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	// Allow up to five seconds to check the connection.
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	// Ping PostgreSQL to make sure it is reachable.
	if err := sqlDB.PingContext(ctx); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("ping PostgreSQL: %w", err)
	}

	// Return the database handle after a successful connection.
	return db, nil
}
