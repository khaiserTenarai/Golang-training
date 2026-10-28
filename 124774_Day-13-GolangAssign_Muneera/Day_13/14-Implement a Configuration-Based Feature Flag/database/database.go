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

// Connect opens and checks the PostgreSQL connection.
func Connect(
	cfg *config.Config,
	secrets *config.Secrets,
) (*gorm.DB, error) {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.DBUser, secrets.DBPassword),
		Host:   net.JoinHostPort(cfg.DBHost, cfg.DBPort),
		Path:   cfg.DBName,
	}

	query := dsn.Query()
	query.Set("sslmode", cfg.DBSSLMode)
	dsn.RawQuery = query.Encode()

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get SQL database: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// Check the database connection with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	// Create or update the employee table.
	if err := db.WithContext(ctx).AutoMigrate(&model.Employee{}); err != nil {
		_ = sqlDB.Close()
		return nil, fmt.Errorf("migrate database: %w", err)
	}

	return db, nil
}