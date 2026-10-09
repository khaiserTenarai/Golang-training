package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"bank-account-system/model"
)

type PostgresAccountRepository struct {
	db *pgxpool.Pool
}

func NewPostgresAccountRepository(db *pgxpool.Pool) *PostgresAccountRepository {
	return &PostgresAccountRepository{db: db}
}

func (r *PostgresAccountRepository) CreateAccount(account model.Account) error {
	query := `
        INSERT INTO accounts (account_number, customer_name, balance)
        VALUES ($1, $2, $3)
    `
	_, err := r.db.Exec(
		context.Background(),
		query,
		account.AccountNumber,
		account.CustomerName,
		account.Balance,
	)
	return err
}

func (r *PostgresAccountRepository) GetAccountByNumber(accountNumber string) (model.Account, error) {
	query := `
        SELECT id, account_number, customer_name, balance, created_at
        FROM accounts
        WHERE account_number = $1
    `

	var account model.Account

	err := r.db.QueryRow(
		context.Background(),
		query,
		accountNumber,
	).Scan(
		&account.ID,
		&account.AccountNumber,
		&account.CustomerName,
		&account.Balance,
		&account.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return model.Account{}, fmt.Errorf(
			"account %s not found",
			accountNumber,
		)
	}

	return account, err
}

func (r *PostgresAccountRepository) UpdateBalance(
	accountID int,
	balance float64,
) error {
	query := `
        UPDATE accounts
        SET balance = $1
        WHERE id = $2
    `

	result, err := r.db.Exec(
		context.Background(),
		query,
		balance,
		accountID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("account not found")
	}

	return nil
}

func (r *PostgresAccountRepository) AddTransaction(
	transaction model.Transaction,
) error {
	query := `
        INSERT INTO transactions
        (account_id, transaction_type, amount)
        VALUES ($1, $2, $3)
    `

	_, err := r.db.Exec(
		context.Background(),
		query,
		transaction.AccountID,
		transaction.TransactionType,
		transaction.Amount,
	)

	return err
}

func (r *PostgresAccountRepository) GetTransactions(
	accountID int,
) ([]model.Transaction, error) {
	query := `
        SELECT id, account_id, transaction_type, amount, created_at
        FROM transactions
        WHERE account_id = $1
        ORDER BY id
    `

	rows, err := r.db.Query(
		context.Background(),
		query,
		accountID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []model.Transaction

	for rows.Next() {
		var transaction model.Transaction

		if err := rows.Scan(
			&transaction.ID,
			&transaction.AccountID,
			&transaction.TransactionType,
			&transaction.Amount,
			&transaction.CreatedAt,
		); err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, rows.Err()
}
