// Package repository has all the direct database access for accounts
// and transactions.
package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"question7-bank-account-system/model"
)

// ErrAccountNotFound is returned when an account ID doesn't exist.
var ErrAccountNotFound = errors.New("account not found")

// ErrInsufficientFunds is returned when a withdrawal would take the
// balance below zero.
var ErrInsufficientFunds = errors.New("insufficient funds for this withdrawal")

// ErrInvalidAmount is returned when a deposit/withdrawal amount is
// zero or negative.
var ErrInvalidAmount = errors.New("amount must be greater than zero")

// CreateAccount opens a new account with the given opening balance.
// If the opening balance is greater than zero, an initial DEPOSIT
// transaction is recorded too - both happen in one transaction, so
// you never end up with an account that has a balance but no matching
// history entry (or vice versa).
func CreateAccount(db *sql.DB, holderName string, openingBalance float64) (int, error) {
	if openingBalance < 0 {
		return 0, fmt.Errorf("create account failed: %w", ErrInvalidAmount)
	}

	tx, err := db.Begin()
	if err != nil {
		return 0, fmt.Errorf("could not start transaction: %w", err)
	}
	defer tx.Rollback()

	var accountID int
	err = tx.QueryRow(
		`INSERT INTO account (account_holder_name, balance) VALUES ($1, $2) RETURNING id`,
		holderName, openingBalance,
	).Scan(&accountID)
	if err != nil {
		return 0, fmt.Errorf("create account failed: %w", err)
	}

	if openingBalance > 0 {
		_, err = tx.Exec(
			`INSERT INTO transaction (account_id, type, amount, balance_after) VALUES ($1, 'DEPOSIT', $2, $3)`,
			accountID, openingBalance, openingBalance,
		)
		if err != nil {
			return 0, fmt.Errorf("create account failed: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("could not commit transaction: %w", err)
	}
	return accountID, nil
}

// GetAccount fetches one account by ID.
func GetAccount(db *sql.DB, id int) (model.Account, error) {
	var a model.Account
	query := `SELECT id, account_holder_name, balance, created_at FROM account WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&a.ID, &a.AccountHolderName, &a.Balance, &a.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Account{}, fmt.Errorf("get account failed: %w", ErrAccountNotFound)
		}
		return model.Account{}, fmt.Errorf("get account failed: %w", err)
	}
	return a, nil
}

// GetBalance is a small convenience wrapper for a balance enquiry.
func GetBalance(db *sql.DB, id int) (float64, error) {
	account, err := GetAccount(db, id)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}

// Deposit adds `amount` to an account's balance and records a DEPOSIT
// transaction, both inside a single database transaction.
func Deposit(db *sql.DB, accountID int, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit failed: %w", ErrInvalidAmount)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("could not start transaction: %w", err)
	}
	defer tx.Rollback()

	var currentBalance float64
	// "FOR UPDATE" locks this account's row until the transaction
	// ends, so two deposits/withdrawals on the same account can't
	// race each other and read a stale balance.
	err = tx.QueryRow(`SELECT balance FROM account WHERE id = $1 FOR UPDATE`, accountID).Scan(&currentBalance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("deposit failed: %w", ErrAccountNotFound)
		}
		return fmt.Errorf("deposit failed: %w", err)
	}

	newBalance := currentBalance + amount

	if _, err := tx.Exec(`UPDATE account SET balance = $1 WHERE id = $2`, newBalance, accountID); err != nil {
		return fmt.Errorf("deposit failed: %w", err)
	}

	if _, err := tx.Exec(
		`INSERT INTO transaction (account_id, type, amount, balance_after) VALUES ($1, 'DEPOSIT', $2, $3)`,
		accountID, amount, newBalance,
	); err != nil {
		return fmt.Errorf("record deposit transaction failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}
	return nil
}

// Withdraw removes `amount` from an account's balance and records a
// WITHDRAWAL transaction, both inside a single database transaction.
// It refuses the withdrawal (rolling everything back) if the balance
// would go negative.
func Withdraw(db *sql.DB, accountID int, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal failed: %w", ErrInvalidAmount)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("could not start transaction: %w", err)
	}
	defer tx.Rollback()

	var currentBalance float64
	err = tx.QueryRow(`SELECT balance FROM account WHERE id = $1 FOR UPDATE`, accountID).Scan(&currentBalance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("withdrawal failed: %w", ErrAccountNotFound)
		}
		return fmt.Errorf("withdrawal failed: %w", err)
	}

	if currentBalance < amount {
		return fmt.Errorf("withdrawal failed: %w", ErrInsufficientFunds)
	}

	newBalance := currentBalance - amount

	if _, err := tx.Exec(`UPDATE account SET balance = $1 WHERE id = $2`, newBalance, accountID); err != nil {
		return fmt.Errorf("withdrawal failed: %w", err)
	}

	if _, err := tx.Exec(
		`INSERT INTO transaction (account_id, type, amount, balance_after) VALUES ($1, 'WITHDRAWAL', $2, $3)`,
		accountID, amount, newBalance,
	); err != nil {
		return fmt.Errorf("record withdrawal transaction failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}
	return nil
}

// GetTransactionHistory returns every deposit/withdrawal for one
// account, most recent first.
func GetTransactionHistory(db *sql.DB, accountID int) ([]model.Transaction, error) {
	query := `
		SELECT id, account_id, type, amount, balance_after, created_at
		FROM transaction
		WHERE account_id = $1
		ORDER BY created_at DESC`
	rows, err := db.Query(query, accountID)
	if err != nil {
		return nil, fmt.Errorf("get transaction history failed: %w", err)
	}
	defer rows.Close()

	var history []model.Transaction
	for rows.Next() {
		var t model.Transaction
		if err := rows.Scan(&t.ID, &t.AccountID, &t.Type, &t.Amount, &t.BalanceAfter, &t.CreatedAt); err != nil {
			return nil, fmt.Errorf("get transaction history failed: %w", err)
		}
		history = append(history, t)
	}
	return history, rows.Err()
}
