package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	connString := "postgres://postgres:password@localhost:5432/bankdb"
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Println("Connection failed:", err)
		os.Exit(1)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)
	ctx := context.Background()

	for {
		fmt.Println("\n1. Create Account")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Balance Enquiry")
		fmt.Println("5. Transaction History")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		switch readLine(reader) {
		case "1":
			createAccount(ctx, db, reader)
		case "2":
			deposit(ctx, db, reader)
		case "3":
			withdraw(ctx, db, reader)
		case "4":
			balanceEnquiry(ctx, db, reader)
		case "5":
			transactionHistory(ctx, db, reader)
		case "6":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func createAccount(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Account holder name: ")
	name := readLine(reader)

	fmt.Print("Opening balance: ")
	balance, _ := strconv.ParseFloat(readLine(reader), 64)

	var id int
	err := db.QueryRow(ctx, `INSERT INTO accounts (name, balance) VALUES ($1, $2) RETURNING id`, name, balance).Scan(&id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Created account with ID:", id)
}

func deposit(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Account ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	fmt.Print("Amount: ")
	amount, _ := strconv.ParseFloat(readLine(reader), 64)

	tx, err := db.Begin(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer tx.Rollback(ctx)

	result, err := tx.Exec(ctx, `UPDATE accounts SET balance = balance + $1 WHERE id = $2`, amount, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if result.RowsAffected() == 0 {
		fmt.Println("Account not found")
		return
	}

	_, err = tx.Exec(ctx, `INSERT INTO transactions (account_id, type, amount) VALUES ($1, 'deposit', $2)`, id, amount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Deposit successful")
}

func withdraw(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Account ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	fmt.Print("Amount: ")
	amount, _ := strconv.ParseFloat(readLine(reader), 64)

	tx, err := db.Begin(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer tx.Rollback(ctx)

	var balance float64
	err = tx.QueryRow(ctx, `SELECT balance FROM accounts WHERE id = $1 FOR UPDATE`, id).Scan(&balance)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if balance < amount {
		fmt.Println("Insufficient balance")
		return
	}

	_, err = tx.Exec(ctx, `UPDATE accounts SET balance = balance - $1 WHERE id = $2`, amount, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, err = tx.Exec(ctx, `INSERT INTO transactions (account_id, type, amount) VALUES ($1, 'withdraw', $2)`, id, amount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := tx.Commit(ctx); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Withdrawal successful")
}

func balanceEnquiry(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Account ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	var balance float64
	err := db.QueryRow(ctx, `SELECT balance FROM accounts WHERE id = $1`, id).Scan(&balance)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Balance: %.2f\n", balance)
}

func transactionHistory(ctx context.Context, db *pgxpool.Pool, reader *bufio.Reader) {
	fmt.Print("Account ID: ")
	id, _ := strconv.Atoi(readLine(reader))

	rows, err := db.Query(ctx, `SELECT type, amount, created_at FROM transactions WHERE account_id = $1 ORDER BY created_at`, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		var txType string
		var amount float64
		var createdAt string
		if err := rows.Scan(&txType, &amount, &createdAt); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%s | %s | %.2f\n", createdAt, txType, amount)
		found = true
	}
	if !found {
		fmt.Println("No transactions found")
	}
}
