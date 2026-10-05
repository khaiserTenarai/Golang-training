package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrAccountNotFound = errors.New("account not found")
var ErrInsufficientBalance = errors.New("insufficient balance")

type AccountRepository interface {
	Create(ctx context.Context, name string, balance float64) error
	List(ctx context.Context) error
	Transfer(ctx context.Context, fromID, toID int64, amount float64) error
}

type PostgresAccountRepository struct{ db *pgxpool.Pool }

func NewPostgresAccountRepository(db *pgxpool.Pool) *PostgresAccountRepository {
	return &PostgresAccountRepository{db: db}
}

func (r *PostgresAccountRepository) Create(ctx context.Context, name string, balance float64) error {
	_, err := r.db.Exec(ctx, `INSERT INTO accounts(name,balance) VALUES($1,$2)`, name, balance)
	return err
}

func (r *PostgresAccountRepository) List(ctx context.Context) error {
	rows, err := r.db.Query(ctx, `SELECT id,name,balance FROM accounts ORDER BY id`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var name string
		var balance float64
		if err := rows.Scan(&id, &name, &balance); err != nil {
			return err
		}
		fmtAccount(id, name, balance)
	}
	return rows.Err()
}

func fmtAccount(id int64, name string, balance float64) {
	fmt.Printf("%d | %s | %.2f\n", id, name, balance)
}

func (r *PostgresAccountRepository) Transfer(ctx context.Context, fromID, toID int64, amount float64) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var balance float64
	if err := tx.QueryRow(ctx, `SELECT balance FROM accounts WHERE id=$1 FOR UPDATE`, fromID).Scan(&balance); err != nil {
		return ErrAccountNotFound
	}
	var destinationID int64
	if err := tx.QueryRow(ctx, `SELECT id FROM accounts WHERE id=$1 FOR UPDATE`, toID).Scan(&destinationID); err != nil {
		return ErrAccountNotFound
	}
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if balance < amount {
		return ErrInsufficientBalance
	}
	if _, err := tx.Exec(ctx, `UPDATE accounts SET balance=balance-$1 WHERE id=$2`, amount, fromID); err != nil {
		return err
	}
	if _, err := tx.Exec(ctx, `UPDATE accounts SET balance=balance+$1 WHERE id=$2`, amount, toID); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
