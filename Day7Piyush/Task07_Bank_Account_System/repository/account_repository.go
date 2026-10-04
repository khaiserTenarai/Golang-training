package repository

import (
	"database/sql"
	"fmt"
	"task07_bank_account_system/models"
)

type AccountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (r *AccountRepository) CreateAccount(acc models.Account) (int, error) {
	var id int
	err := r.DB.QueryRow(
		`INSERT INTO bank_accounts (holder_name, account_type, balance) VALUES ($1, $2, $3) RETURNING id`,
		acc.HolderName, acc.AccountType, acc.Balance).Scan(&id)
	return id, err
}

func (r *AccountRepository) GetByID(id int) (models.Account, error) {
	var acc models.Account
	err := r.DB.QueryRow(
		`SELECT id, holder_name, account_type, balance, created_at FROM bank_accounts WHERE id=$1`, id).
		Scan(&acc.ID, &acc.HolderName, &acc.AccountType, &acc.Balance, &acc.CreatedAt)
	return acc, err
}

func (r *AccountRepository) GetAll() ([]models.Account, error) {
	rows, err := r.DB.Query(`SELECT id, holder_name, account_type, balance, created_at FROM bank_accounts ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var accounts []models.Account
	for rows.Next() {
		var acc models.Account
		if err := rows.Scan(&acc.ID, &acc.HolderName, &acc.AccountType, &acc.Balance, &acc.CreatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}
	return accounts, rows.Err()
}

func (r *AccountRepository) Deposit(accountID int, amount float64, description string) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	var balance float64
	err = tx.QueryRow(`SELECT balance FROM bank_accounts WHERE id=$1 FOR UPDATE`, accountID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("account not found: %w", err)
	}

	newBalance := balance + amount
	_, err = tx.Exec(`UPDATE bank_accounts SET balance=$1 WHERE id=$2`, newBalance, accountID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("update balance: %w", err)
	}

	_, err = tx.Exec(
		`INSERT INTO bank_transactions (account_id, txn_type, amount, balance_after, description) VALUES ($1, 'DEPOSIT', $2, $3, $4)`,
		accountID, amount, newBalance, description)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("record transaction: %w", err)
	}

	return tx.Commit()
}

func (r *AccountRepository) Withdraw(accountID int, amount float64, description string) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	var balance float64
	err = tx.QueryRow(`SELECT balance FROM bank_accounts WHERE id=$1 FOR UPDATE`, accountID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("account not found: %w", err)
	}

	if balance < amount {
		tx.Rollback()
		return fmt.Errorf("insufficient balance: available %.2f, requested %.2f", balance, amount)
	}

	newBalance := balance - amount
	_, err = tx.Exec(`UPDATE bank_accounts SET balance=$1 WHERE id=$2`, newBalance, accountID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("update balance: %w", err)
	}

	_, err = tx.Exec(
		`INSERT INTO bank_transactions (account_id, txn_type, amount, balance_after, description) VALUES ($1, 'WITHDRAWAL', $2, $3, $4)`,
		accountID, amount, newBalance, description)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("record transaction: %w", err)
	}

	return tx.Commit()
}

func (r *AccountRepository) GetTransactionHistory(accountID int) ([]models.Transaction, error) {
	rows, err := r.DB.Query(
		`SELECT id, account_id, txn_type, amount, balance_after, description, created_at FROM bank_transactions WHERE account_id=$1 ORDER BY created_at DESC`,
		accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var txns []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.AccountID, &t.TxnType, &t.Amount, &t.BalanceAfter, &t.Description, &t.CreatedAt); err != nil {
			return nil, err
		}
		txns = append(txns, t)
	}
	return txns, rows.Err()
}
