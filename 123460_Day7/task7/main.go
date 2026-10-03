package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/jackc/pgx/v5"
)

var reader = bufio.NewReader(os.Stdin)

func readInput(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	connString := "postgres://postgres:pgadmin@localhost:5432/employee_db"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("PostgreSQL connected successfully")

	for {
		fmt.Println("\n==============================")
		fmt.Println("      BANK ACCOUNT SYSTEM")
		fmt.Println("==============================")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Check Balance")
		fmt.Println("5. Transaction History")
		fmt.Println("6. Exit")

		choiceStr := readInput("Choose an option: ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 6.")
			continue
		}

		switch choice {
		case 1:
			createAccount(conn)
		case 2:
			processTransaction(conn, "DEPOSIT")
		case 3:
			processTransaction(conn, "WITHDRAWAL")
		case 4:
			checkBalance(conn)
		case 5:
			viewTransactionHistory(conn)
		case 6:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please select between 1 and 6.")
		}
	}
}

func createAccount(conn *pgx.Conn) {
	fmt.Println("\n----- CREATE ACCOUNT -----")
	name := readInput("Customer Name: ")

	initialDepositStr := readInput("Initial Deposit Amount: ")
	initialDeposit, err := strconv.ParseFloat(initialDepositStr, 64)
	if err != nil || initialDeposit < 0 {
		fmt.Println("Invalid amount. Starting balance will be $0.00")
		initialDeposit = 0.00
	}

	ctx := context.Background()
	
	// Use a transaction to create account and record initial deposit history if > 0
	tx, err := conn.Begin(ctx)
	if err != nil {
		fmt.Println("Failed to start transaction:", err)
		return
	}
	defer tx.Rollback(ctx)

	var accountID int
	err = tx.QueryRow(
		ctx,
		`INSERT INTO accounts (customer_name, balance) VALUES ($1, $2) RETURNING id`,
		name, initialDeposit,
	).Scan(&accountID)

	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}

	if initialDeposit > 0 {
		_, err = tx.Exec(
			ctx,
			`INSERT INTO transactions (account_id, transaction_type, amount) VALUES ($1, 'DEPOSIT', $2)`,
			accountID, initialDeposit,
		)
		if err != nil {
			fmt.Println("Error recording initial deposit:", err)
			return
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		fmt.Println("Failed to commit transaction:", err)
		return
	}

	fmt.Printf("Account created successfully. Account ID: %d | Balance: $%.2f\n", accountID, initialDeposit)
}

func processTransaction(conn *pgx.Conn, txType string) {
	fmt.Printf("\n----- %s -----\n", txType)
	idStr := readInput("Enter Account ID: ")
	accountID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid Account ID.")
		return
	}

	amountStr := readInput("Enter Amount: ")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		fmt.Println("Invalid amount. Must be greater than 0.")
		return
	}

	ctx := context.Background()

	// Start Database Transaction
	tx, err := conn.Begin(ctx)
	if err != nil {
		fmt.Println("Failed to start transaction:", err)
		return
	}
	defer tx.Rollback(ctx)

	// Fetch current balance
	var currentBalance float64
	err = tx.QueryRow(ctx, `SELECT balance FROM accounts WHERE id = $1`, accountID).Scan(&currentBalance)
	if err != nil {
		fmt.Println("Error: Account not found.")
		return
	}

	if txType == "WITHDRAWAL" && currentBalance < amount {
		fmt.Printf("Insufficient funds. Current Balance: $%.2f\n", currentBalance)
		return
	}

	// Update Balance
	var updateQuery string
	if txType == "DEPOSIT" {
		updateQuery = `UPDATE accounts SET balance = balance + $1 WHERE id = $2 RETURNING balance`
	} else {
		updateQuery = `UPDATE accounts SET balance = balance - $1 WHERE id = $2 RETURNING balance`
	}

	var newBalance float64
	err = tx.QueryRow(ctx, updateQuery, amount, accountID).Scan(&newBalance)
	if err != nil {
		fmt.Println("Error updating account balance:", err)
		return
	}

	// Record Transaction History
	_, err = tx.Exec(
		ctx,
		`INSERT INTO transactions (account_id, transaction_type, amount) VALUES ($1, $2, $3)`,
		accountID, txType, amount,
	)
	if err != nil {
		fmt.Println("Error recording transaction history:", err)
		return
	}

	// Commit Transaction
	err = tx.Commit(ctx)
	if err != nil {
		fmt.Println("Failed to commit transaction:", err)
		return
	}

	fmt.Printf("%s successful! New Balance: $%.2f\n", txType, newBalance)
}

func checkBalance(conn *pgx.Conn) {
	fmt.Println("\n----- BALANCE ENQUIRY -----")
	idStr := readInput("Enter Account ID: ")
	accountID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid Account ID.")
		return
	}

	var name string
	var balance float64

	err = conn.QueryRow(
		context.Background(),
		`SELECT customer_name, balance FROM accounts WHERE id = $1`,
		accountID,
	).Scan(&name, &balance)

	if err != nil {
		fmt.Println("Error: Account not found.")
		return
	}

	fmt.Printf("Account: %s | Current Balance: $%.2f\n", name, balance)
}

func viewTransactionHistory(conn *pgx.Conn) {
	fmt.Println("\n----- TRANSACTION HISTORY -----")
	idStr := readInput("Enter Account ID: ")
	accountID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid Account ID.")
		return
	}

	rows, err := conn.Query(
		context.Background(),
		`SELECT transaction_type, amount, timestamp 
		 FROM transactions 
		 WHERE account_id = $1 
		 ORDER BY timestamp DESC`,
		accountID,
	)
	if err != nil {
		fmt.Println("Error fetching transactions:", err)
		return
	}
	defer rows.Close()

	fmt.Printf("\n--- History for Account ID: %d ---\n", accountID)
	count := 0
	for rows.Next() {
		var txType string
		var amount float64
		var timestamp time.Time 

		if err := rows.Scan(&txType, &amount, &timestamp); err != nil {
			fmt.Println("Error reading row:", err) // Print the error instead of failing silently
			continue
		}
		count++
		
		// Format the timestamp to a readable string format
		fmt.Printf("[%s] Type: %-10s | Amount: $%.2f\n", timestamp.Format("2006-01-02 15:04:05"), txType, amount)
	}

	if count == 0 {
		fmt.Println("No transactions found for this account.")
	}
}