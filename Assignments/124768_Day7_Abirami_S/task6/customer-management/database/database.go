package database

import (
	"context"
	"log"

	"customer-management/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() *pgxpool.Pool {
	cfg := config.Load()

	db, err := pgxpool.New(context.Background(), cfg.DatabaseURL())
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal("Database ping failed:", err)
	}

	log.Println("Database connected successfully")

	return db
}
