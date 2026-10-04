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
	accQuery := `CREATE TABLE IF NOT EXISTS bank_accounts (
		id SERIAL PRIMARY KEY,
		holder_name VARCHAR(100) NOT NULL,
		account_type VARCHAR(20) DEFAULT 'savings',
		balance DECIMAL(15,2) DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	if _, err := db.Exec(accQuery); err != nil {
		log.Fatal("Failed to create bank_accounts table:", err)
	}

	txnQuery := `CREATE TABLE IF NOT EXISTS bank_transactions (
		id SERIAL PRIMARY KEY,
		account_id INT REFERENCES bank_accounts(id) ON DELETE CASCADE,
		txn_type VARCHAR(20) NOT NULL,
		amount DECIMAL(15,2) NOT NULL,
		balance_after DECIMAL(15,2) NOT NULL,
		description VARCHAR(200),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	if _, err := db.Exec(txnQuery); err != nil {
		log.Fatal("Failed to create bank_transactions table:", err)
	}
	fmt.Println("Tables 'bank_accounts' and 'bank_transactions' are ready.")
}
