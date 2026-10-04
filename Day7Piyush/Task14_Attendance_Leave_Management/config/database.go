package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "Piyushgoyal@2005"
	DBName   = "day7db"
)

func ConnectDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping:", err)
	}
	fmt.Println("Connected to PostgreSQL successfully!")
	return db
}

func CreateTables(db *sql.DB) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS att_employees (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			department VARCHAR(50),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS attendance (
			id SERIAL PRIMARY KEY,
			employee_id INT REFERENCES att_employees(id) ON DELETE CASCADE,
			check_in TIMESTAMP,
			check_out TIMESTAMP,
			date DATE DEFAULT CURRENT_DATE,
			status VARCHAR(20) DEFAULT 'present'
		)`,
		`CREATE TABLE IF NOT EXISTS leave_requests (
			id SERIAL PRIMARY KEY,
			employee_id INT REFERENCES att_employees(id) ON DELETE CASCADE,
			leave_type VARCHAR(30) NOT NULL,
			start_date DATE NOT NULL,
			end_date DATE NOT NULL,
			reason VARCHAR(300),
			status VARCHAR(20) DEFAULT 'pending',
			reviewed_by VARCHAR(100),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
	}
	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			log.Fatal("Table creation error:", err)
		}
	}
	fmt.Println("Attendance & Leave tables are ready.")
}
