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
	accQuery := `CREATE TABLE IF NOT EXISTS transfer_accounts (
		id SERIAL PRIMARY KEY,
		holder_name VARCHAR(100) NOT NULL,
		balance DECIMAL(15,2) DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	if _, err := db.Exec(accQuery); err != nil {
		log.Fatal("Failed to create transfer_accounts table:", err)
	}

	txnQuery := `CREATE TABLE IF NOT EXISTS transfer_log (
		id SERIAL PRIMARY KEY,
		from_account_id INT REFERENCES transfer_accounts(id),
		to_account_id INT REFERENCES transfer_accounts(id),
		amount DECIMAL(15,2) NOT NULL,
		status VARCHAR(20) DEFAULT 'SUCCESS',
		description VARCHAR(200),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	if _, err := db.Exec(txnQuery); err != nil {
		log.Fatal("Failed to create transfer_log table:", err)
	}
	fmt.Println("Tables 'transfer_accounts' and 'transfer_log' are ready.")
}
