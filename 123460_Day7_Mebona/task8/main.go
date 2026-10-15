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
		fmt.Println("      MONEY TRANSFER SYSTEM")
		fmt.Println("==============================")
		fmt.Println("1. Create Account (Setup)")
		fmt.Println("2. Transfer Money")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Transaction History")
		fmt.Println("5. Exit")

		choiceStr := readInput("Choose an option: ")
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 5.")
			continue
		}

		switch choice {
		case 1:
			createAccount(conn)
		case 2:
			transferMoney(conn)
		case 3:
			checkBalance(conn)
		case 4:
			viewTransactionHistory(conn)
		case 5:
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please select between 1 and 5.")
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
			`INSERT INTO transactions (account_id, transaction_type, amount) VALUES ($1, 'INITIAL_FUND', $2)`,
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

func transferMoney(conn *pgx.Conn) {
	fmt.Println("\n----- TRANSFER MONEY -----")
	
	senderIDStr := readInput("Enter Sender Account ID: ")
	senderID, err := strconv.Atoi(senderIDStr)
	if err != nil {
		fmt.Println("Invalid Sender Account ID.")
		return
	}

	receiverIDStr := readInput("Enter Receiver Account ID: ")
	receiverID, err := strconv.Atoi(receiverIDStr)
	if err != nil {
		fmt.Println("Invalid Receiver Account ID.")
		return
	}

	if senderID == receiverID {
		fmt.Println("Cannot transfer money to the same account.")
		return
	}

	amountStr := readInput("Enter Transfer Amount: ")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		fmt.Println("Invalid transfer amount. Must be greater than 0.")
		return
	}

	ctx := context.Background()

	// ========================================================
	// 1. BEGIN TRANSACTION
	// ========================================================
	tx, err := conn.Begin(ctx)
	if err != nil {
		fmt.Println("Failed to start transaction:", err)
		return
	}
	
	// Ensure rollback happens if the function exits early due to an error. 
	// If tx.Commit() succeeds at the end, this rollback becomes a no-op.
	defer tx.Rollback(ctx)

	// 2. Check Sender Balance
	var senderBalance float64
	err = tx.QueryRow(ctx, `SELECT balance FROM accounts WHERE id = $1`, senderID).Scan(&senderBalance)
	if err != nil {
		fmt.Println("Error: Sender account not found.")
		return // Triggers rollback automatically
	}

	if senderBalance < amount {
		fmt.Printf("Insufficient funds in sender account. Current Balance: $%.2f\n", senderBalance)
		return // Triggers rollback automatically
	}

	// 3. Verify Receiver Exists
	var receiverIDCheck int
	err = tx.QueryRow(ctx, `SELECT id FROM accounts WHERE id = $1`, receiverID).Scan(&receiverIDCheck)
	if err != nil {
		fmt.Println("Error: Receiver account not found.")
		return // Triggers rollback automatically
	}

	// 4. Deduct from Sender
	var newSenderBalance float64
	err = tx.QueryRow(ctx, `UPDATE accounts SET balance = balance - $1 WHERE id = $2 RETURNING balance`, amount, senderID).Scan(&newSenderBalance)
	if err != nil {
		fmt.Println("Error deducting from sender:", err)
		return
	}

	// 5. Add to Receiver
	_, err = tx.Exec(ctx, `UPDATE accounts SET balance = balance + $1 WHERE id = $2`, amount, receiverID)
	if err != nil {
		fmt.Println("Error adding to receiver:", err)
		return
	}

	// 6. Record Sender Transaction History (Money Out)
	_, err = tx.Exec(ctx, `INSERT INTO transactions (account_id, transaction_type, amount) VALUES ($1, 'TRANSFER_OUT', $2)`, senderID, amount)
	if err != nil {
		fmt.Println("Error recording sender transaction:", err)
		return
	}

	// 7. Record Receiver Transaction History (Money In)
	_, err = tx.Exec(ctx, `INSERT INTO transactions (account_id, transaction_type, amount) VALUES ($1, 'TRANSFER_IN', $2)`, receiverID, amount)
	if err != nil {
		fmt.Println("Error recording receiver transaction:", err)
		return
	}

	// ========================================================
	// 8. COMMIT TRANSACTION (Only executes if everything above passed)
	// ========================================================
	err = tx.Commit(ctx)
	if err != nil {
		fmt.Println("Failed to commit transaction:", err)
		return
	}

	fmt.Printf("Transfer of $%.2f successful! Sender New Balance: $%.2f\n", amount, newSenderBalance)
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
			fmt.Println("Error reading row:", err)
			continue
		}
		count++
		
		fmt.Printf("[%s] Type: %-12s | Amount: $%.2f\n", timestamp.Format("2006-01-02 15:04:05"), txType, amount)
	}

	if count == 0 {
		fmt.Println("No transactions found for this account.")
	}
}