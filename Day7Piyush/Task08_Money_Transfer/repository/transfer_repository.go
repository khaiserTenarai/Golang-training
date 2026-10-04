package repository

import (
	"database/sql"
	"fmt"
	"task08_money_transfer/models"
)

type TransferRepository struct {
	DB *sql.DB
}

func NewTransferRepository(db *sql.DB) *TransferRepository {
	return &TransferRepository{DB: db}
}

func (r *TransferRepository) CreateAccount(acc models.Account) (int, error) {
	var id int
	err := r.DB.QueryRow(
		`INSERT INTO transfer_accounts (holder_name, balance) VALUES ($1, $2) RETURNING id`,
		acc.HolderName, acc.Balance).Scan(&id)
	return id, err
}

func (r *TransferRepository) GetAll() ([]models.Account, error) {
	rows, err := r.DB.Query(`SELECT id, holder_name, balance, created_at FROM transfer_accounts ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var accounts []models.Account
	for rows.Next() {
		var acc models.Account
		if err := rows.Scan(&acc.ID, &acc.HolderName, &acc.Balance, &acc.CreatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}
	return accounts, rows.Err()
}

func (r *TransferRepository) GetByID(id int) (models.Account, error) {
	var acc models.Account
	err := r.DB.QueryRow(
		`SELECT id, holder_name, balance, created_at FROM transfer_accounts WHERE id=$1`, id).
		Scan(&acc.ID, &acc.HolderName, &acc.Balance, &acc.CreatedAt)
	return acc, err
}

// Transfer uses a PostgreSQL transaction with commit and rollback
func (r *TransferRepository) Transfer(fromID, toID int, amount float64, description string) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	// Lock both accounts using FOR UPDATE to prevent race conditions
	var fromBalance float64
	err = tx.QueryRow(`SELECT balance FROM transfer_accounts WHERE id=$1 FOR UPDATE`, fromID).Scan(&fromBalance)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("source account not found: %w", err)
	}

	var toBalance float64
	err = tx.QueryRow(`SELECT balance FROM transfer_accounts WHERE id=$1 FOR UPDATE`, toID).Scan(&toBalance)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("destination account not found: %w", err)
	}

	// Check sufficient balance
	if fromBalance < amount {
		tx.Rollback()
		// Log failed transfer
		r.DB.Exec(
			`INSERT INTO transfer_log (from_account_id, to_account_id, amount, status, description) VALUES ($1, $2, $3, 'FAILED', $4)`,
			fromID, toID, amount, "Insufficient balance - rolled back")
		return fmt.Errorf("insufficient balance: available %.2f, requested %.2f (ROLLED BACK)", fromBalance, amount)
	}

	// Debit from source
	_, err = tx.Exec(`UPDATE transfer_accounts SET balance = balance - $1 WHERE id = $2`, amount, fromID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("debit failed: %w", err)
	}

	// Credit to destination
	_, err = tx.Exec(`UPDATE transfer_accounts SET balance = balance + $1 WHERE id = $2`, amount, toID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("credit failed: %w", err)
	}

	// Log the successful transfer
	_, err = tx.Exec(
		`INSERT INTO transfer_log (from_account_id, to_account_id, amount, status, description) VALUES ($1, $2, $3, 'SUCCESS', $4)`,
		fromID, toID, amount, description)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("log transfer: %w", err)
	}

	// COMMIT the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	fmt.Println("[TX] Transaction committed successfully.")
	return nil
}

func (r *TransferRepository) GetTransferLogs() ([]models.TransferLog, error) {
	query := `SELECT t.id, t.from_account_id, f.holder_name, t.to_account_id, toa.holder_name, 
			t.amount, t.status, COALESCE(t.description,''), t.created_at
		FROM transfer_log t
		JOIN transfer_accounts f ON t.from_account_id = f.id
		JOIN transfer_accounts toa ON t.to_account_id = toa.id
		ORDER BY t.created_at DESC`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var logs []models.TransferLog
	for rows.Next() {
		var l models.TransferLog
		if err := rows.Scan(&l.ID, &l.FromAccountID, &l.FromName, &l.ToAccountID, &l.ToName,
			&l.Amount, &l.Status, &l.Description, &l.CreatedAt); err != nil {
			return nil, err
		}
		logs = append(logs, l)
	}
	return logs, rows.Err()
}
