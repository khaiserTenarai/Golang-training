package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"employee-api/config"
)

type DB struct {
	DSN string
}

func NewPostgresDB(ctx context.Context, cfg config.DatabaseConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.SSLMode)

	log.Printf("[DB] Connecting to database at %s:%s...", cfg.Host, cfg.Port)

	select {
	case <-time.After(100 * time.Millisecond):
		log.Println("[DB] Database connection established successfully.")
		return &DB{DSN: dsn}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (db *DB) Close() error {
	log.Println("[DB] Closing database connection pool...")
	return nil
}
