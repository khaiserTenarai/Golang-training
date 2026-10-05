// Package db handles the PostgreSQL connection for this project.
package db

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
	// "github.com/lib/pq"
)

// Connect opens a connection pool to Postgres using environment
// variables (falling back to sensible local defaults), and checks the
// connection actually works before returning.
func Connect() (*sql.DB, error) {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "Info@131")
	dbName := getEnv("DB_NAME", "product_inventory_db")
	sslMode := getEnv("DB_SSLMODE", "disable")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbName, sslMode,
	)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not open connection: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	return conn, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
