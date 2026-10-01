// Bank Account System - an interactive CLI backed by PostgreSQL.
//
// Before running: apply sql/schema.sql to your database, then
//   go mod tidy
//   go run .
package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"question7-bank-account-system/db"
	"question7-bank-account-system/repository"
)

var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func readInt(prompt string) int {
	for {
		text := readLine(prompt)
		value, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Please enter a whole number.")
			continue
		}
		return value
	}
}

func readFloat(prompt string) float64 {
	for {
		text := readLine(prompt)
		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}
		return value
	}
}

func main() {
	conn, err := db.Connect()
	if err != nil {
		fmt.Println("Could not connect to the database:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to Postgres successfully.")

	for {
		fmt.Println("\n--- Bank Account System ---")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Balance Enquiry")
		fmt.Println("5. Transaction History")
		fmt.Println("6. Exit")

		switch readLine("Enter your choice: ") {
		case "1":
			createAccount(conn)
		case "2":
			deposit(conn)
		case "3":
			withdraw(conn)
		case "4":
			balanceEnquiry(conn)
		case "5":
			transactionHistory(conn)
		case "6":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}

func createAccount(conn *sql.DB) {
	holderName := readLine("Account holder name: ")
	openingBalance := readFloat("Opening deposit (0 for none): ")

	id, err := repository.CreateAccount(conn, holderName, openingBalance)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Account created. Your account number is:", id)
}

func deposit(conn *sql.DB) {
	accountID := readInt("Enter account number: ")
	amount := readFloat("Amount to deposit: ")

	err := repository.Deposit(conn, accountID, amount)
	if err != nil {
		printAccountError(err)
		return
	}
	fmt.Println("Deposit successful.")
	printBalance(conn, accountID)
}

func withdraw(conn *sql.DB) {
	accountID := readInt("Enter account number: ")
	amount := readFloat("Amount to withdraw: ")

	err := repository.Withdraw(conn, accountID, amount)
	if err != nil {
		printAccountError(err)
		return
	}
	fmt.Println("Withdrawal successful.")
	printBalance(conn, accountID)
}

func balanceEnquiry(conn *sql.DB) {
	accountID := readInt("Enter account number: ")
	printBalance(conn, accountID)
}

func transactionHistory(conn *sql.DB) {
	accountID := readInt("Enter account number: ")
	history, err := repository.GetTransactionHistory(conn, accountID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(history) == 0 {
		fmt.Println("No transactions yet for this account.")
		return
	}
	fmt.Println("Transaction history (most recent first):")
	for _, t := range history {
		fmt.Printf("  [%d] %-10s amount: %.2f | balance after: %.2f | %s\n",
			t.ID, t.Type, t.Amount, t.BalanceAfter, t.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

func printBalance(conn *sql.DB, accountID int) {
	balance, err := repository.GetBalance(conn, accountID)
	if err != nil {
		printAccountError(err)
		return
	}
	fmt.Printf("Current balance: %.2f\n", balance)
}

func printAccountError(err error) {
	switch {
	case errors.Is(err, repository.ErrAccountNotFound):
		fmt.Println("No account found with that number.")
	case errors.Is(err, repository.ErrInsufficientFunds):
		fmt.Println("Insufficient funds for that withdrawal.")
	case errors.Is(err, repository.ErrInvalidAmount):
		fmt.Println("Amount must be greater than zero.")
	default:
		fmt.Println("Error:", err)
	}
}
