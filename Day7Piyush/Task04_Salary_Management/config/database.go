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
	empQuery := `CREATE TABLE IF NOT EXISTS sal_employees (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		department VARCHAR(50),
		salary DECIMAL(10,2) DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	if _, err := db.Exec(empQuery); err != nil {
		log.Fatal("Failed to create sal_employees table:", err)
	}

	histQuery := `CREATE TABLE IF NOT EXISTS salary_history (
		id SERIAL PRIMARY KEY,
		employee_id INT REFERENCES sal_employees(id) ON DELETE CASCADE,
		old_salary DECIMAL(10,2),
		new_salary DECIMAL(10,2),
		change_reason VARCHAR(200),
		changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	if _, err := db.Exec(histQuery); err != nil {
		log.Fatal("Failed to create salary_history table:", err)
	}
	fmt.Println("Tables 'sal_employees' and 'salary_history' are ready.")
}
