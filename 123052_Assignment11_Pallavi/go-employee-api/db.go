package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=Info@2k26 dbname=assessment2_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS employees (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		role VARCHAR(100) NOT NULL
	);`
	
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return db, nil
}