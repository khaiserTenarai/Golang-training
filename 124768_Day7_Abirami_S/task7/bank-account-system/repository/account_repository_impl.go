package repository

import (
	"context"
	"errors"
	"fmt"

	"bank-account-system/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrAccountNotFound = errors.New("account not found")
var ErrInsufficientBalance = errors.New("insufficient balance")

type AccountRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewAccountRepository(db *pgxpool.Pool) AccountRepository {
	return &AccountRepositoryImpl{
		db: db,
	}
}

func (r *AccountRepositoryImpl) CreateAccount(account model.Account) error {
	fmt.Println("\n----- CREATE ACCOUNT -----")

	query := `
		INSERT INTO accounts (account_number, holder_name, balance)
		VALUES ($1, $2, $3)
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		account.AccountNumber,
		account.HolderName,
		account.Balance,
	)

	if err != nil {
		return err
	}

	fmt.Println("Account created successfully")
	fmt.Println("Rows affected:", result.RowsAffected())

	return nil
}

func (r *AccountRepositoryImpl) GetAccount(id int) (*model.Account, error) {
	fmt.Println("\n----- BALANCE ENQUIRY -----")

	var account model.Account

	query := `
		SELECT id, account_number, holder_name, balance
		FROM accounts
		WHERE id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&account.ID,
		&account.AccountNumber,
		&account.HolderName,
		&account.Balance,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrAccountNotFound
		}

		return nil, err
	}

	return &account, nil
}

func (r *AccountRepositoryImpl) Deposit(id int, amount float64) error {
	fmt.Println("\n----- DEPOSIT -----")

	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}

	result, err := r.db.Exec(
		context.Background(),
		`UPDATE accounts SET balance = balance + $1 WHERE id = $2`,
		amount,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrAccountNotFound
	}

	_, err = r.db.Exec(
		context.Background(),
		`INSERT INTO transactions (account_id, transaction_type, amount)
		 VALUES ($1, $2, $3)`,
		id,
		"DEPOSIT",
		amount,
	)

	if err != nil {
		return err
	}

	fmt.Println("Deposit successful")

	return nil
}

func (r *AccountRepositoryImpl) Withdraw(id int, amount float64) error {
	fmt.Println("\n----- WITHDRAW -----")

	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}

	result, err := r.db.Exec(
		context.Background(),
		`UPDATE accounts
		 SET balance = balance - $1
		 WHERE id = $2 AND balance >= $1`,
		amount,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		account, err := r.GetAccount(id)

		if err != nil {
			return err
		}

		if account.Balance < amount {
			return ErrInsufficientBalance
		}

		return ErrAccountNotFound
	}

	_, err = r.db.Exec(
		context.Background(),
		`INSERT INTO transactions (account_id, transaction_type, amount)
		 VALUES ($1, $2, $3)`,
		id,
		"WITHDRAW",
		amount,
	)

	if err != nil {
		return err
	}

	fmt.Println("Withdrawal successful")

	return nil
}

func (r *AccountRepositoryImpl) GetTransactionHistory(accountID int) ([]model.Transaction, error) {
	fmt.Println("\n----- TRANSACTION HISTORY -----")

	query := `
		SELECT id, account_id, transaction_type, amount, transaction_date
		FROM transactions
		WHERE account_id = $1
		ORDER BY transaction_date DESC
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

		err := rows.Scan(
			&transaction.ID,
			&transaction.AccountID,
			&transaction.TransactionType,
			&transaction.Amount,
			&transaction.TransactionDate,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
