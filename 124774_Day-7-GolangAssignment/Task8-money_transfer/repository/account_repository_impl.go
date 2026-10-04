package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

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

func (r *AccountRepositoryImpl) Transfer(
	fromID int,
	toID int,
	amount float64,
) error {

	ctx := context.Background()

	// Start transaction
	tx, err := r.db.Begin(ctx)

	if err != nil {
		return err
	}

	// Rollback if any operation fails
	defer tx.Rollback(ctx)

	// Check sender balance
	var balance float64

	err = tx.QueryRow(
		ctx,
		`SELECT balance
		 FROM accounts
		 WHERE id = $1
		 FOR UPDATE`,
		fromID,
	).Scan(&balance)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("source account not found")
		}

		return err
	}

	if balance < amount {
		return errors.New("insufficient balance")
	}

	// Deduct from sender
	_, err = tx.Exec(
		ctx,
		`UPDATE accounts
		 SET balance = balance - $1
		 WHERE id = $2`,
		amount,
		fromID,
	)

	if err != nil {
		return err
	}

	// Add to receiver
	result, err := tx.Exec(
		ctx,
		`UPDATE accounts
		 SET balance = balance + $1
		 WHERE id = $2`,
		amount,
		toID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("destination account not found")
	}

	// Commit transaction
	return tx.Commit(ctx)
}
