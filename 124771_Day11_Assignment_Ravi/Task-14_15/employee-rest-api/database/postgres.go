package database

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"employee-rest-api/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(cfg config.Config) (*pgxpool.Pool, error) {
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.DBUser, cfg.DBPassword),
		Host:   fmt.Sprintf("%s:%s", cfg.DBHost, cfg.DBPort),
		Path:   cfg.DBName,
	}
	q := u.Query()
	q.Set("sslmode", cfg.DBSSLMode)
	u.RawQuery = q.Encode()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, u.String())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(ctx); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func CreateSchema(db *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := db.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS employees (
			id BIGSERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(150) UNIQUE NOT NULL,
			department VARCHAR(100) NOT NULL,
			salary NUMERIC(12, 2) NOT NULL CHECK (salary >= 0)
		)
	`)
	return err
}
