package repository

import (
	"context"
	"errors"

	"bankaccount/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("account not found")

type AccountRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewAccountRepository(
	db *pgxpool.Pool,
) AccountRepository {

	return &AccountRepositoryImpl{
		db: db,
	}
}

// CREATE ACCOUNT

func (r *AccountRepositoryImpl) Create(
	account model.Account,
) error {

	query := `
		INSERT INTO accounts
		(account_number, name, balance)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		account.AccountNumber,
		account.Name,
		account.Balance,
	)

	return err
}

// DEPOSIT

func (r *AccountRepositoryImpl) Deposit(
	id int,
	amount float64,
) error {

	tx, err := r.db.Begin(context.Background())

	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	var balance float64

	err = tx.QueryRow(
		context.Background(),
		`SELECT balance FROM accounts WHERE id = $1`,
		id,
	).Scan(&balance)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}

		return err
	}

	newBalance := balance + amount

	_, err = tx.Exec(
		context.Background(),
		`UPDATE accounts SET balance = $1 WHERE id = $2`,
		newBalance,
		id,
	)

	if err != nil {
		return err
	}

	_, err = tx.Exec(
		context.Background(),
		`INSERT INTO transactions
		(account_id, transaction_type, amount, balance_after)
		VALUES ($1, $2, $3, $4)`,
		id,
		"DEPOSIT",
		amount,
		newBalance,
	)

	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

// WITHDRAW

func (r *AccountRepositoryImpl) Withdraw(
	id int,
	amount float64,
) error {

	tx, err := r.db.Begin(context.Background())

	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	var balance float64

	err = tx.QueryRow(
		context.Background(),
		`SELECT balance FROM accounts WHERE id = $1`,
		id,
	).Scan(&balance)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}

		return err
	}

	if balance < amount {
		return errors.New("insufficient balance")
	}

	newBalance := balance - amount

	_, err = tx.Exec(
		context.Background(),
		`UPDATE accounts SET balance = $1 WHERE id = $2`,
		newBalance,
		id,
	)

	if err != nil {
		return err
	}

	_, err = tx.Exec(
		context.Background(),
		`INSERT INTO transactions
		(account_id, transaction_type, amount, balance_after)
		VALUES ($1, $2, $3, $4)`,
		id,
		"WITHDRAW",
		amount,
		newBalance,
	)

	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

// FIND ACCOUNT

func (r *AccountRepositoryImpl) FindByID(
	id int,
) (model.Account, error) {

	var account model.Account

	query := `
		SELECT
			id,
			account_number,
			name,
			balance
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
		&account.Name,
		&account.Balance,
	)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return model.Account{}, ErrNotFound
		}

		return model.Account{}, err
	}

	return account, nil
}

// TRANSACTION HISTORY

func (r *AccountRepositoryImpl) FindTransactions(
	id int,
) ([]model.Transaction, error) {

	query := `
		SELECT
			id,
			account_id,
			transaction_type,
			amount,
			balance_after
		FROM transactions
		WHERE account_id = $1
		ORDER BY id
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
		id,
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
			&transaction.Type,
			&transaction.Amount,
			&transaction.BalanceAfter,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(
			transactions,
			transaction,
		)
	}

	return transactions, rows.Err()
}
