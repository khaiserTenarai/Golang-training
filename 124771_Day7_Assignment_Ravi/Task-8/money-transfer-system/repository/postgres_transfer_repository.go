package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"money-transfer-system/model"
)

type PostgresTransferRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTransferRepository(db *pgxpool.Pool) *PostgresTransferRepository {
	return &PostgresTransferRepository{db: db}
}

func (r *PostgresTransferRepository) GetAccount(
	accountNumber string,
) (model.Account, error) {
	query := `
        SELECT id, account_number, customer_name, balance
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
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return model.Account{}, fmt.Errorf(
			"account %s not found",
			accountNumber,
		)
	}

	return account, err
}

func (r *PostgresTransferRepository) TransferMoney(
	transfer model.Transfer,
) error {
	ctx := context.Background()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("unable to begin transaction: %w", err)
	}

	// Rollback is safe even if Commit succeeds.
	// It also guarantees rollback if any later operation fails.
	defer tx.Rollback(ctx)

	// Lock both accounts before changing balances.
	// Locking by ID order helps reduce deadlock risk when two transfers
	// happen at the same time in opposite directions.
	query := `
        SELECT id, account_number, customer_name, balance
        FROM accounts
        WHERE account_number = $1 OR account_number = $2
        ORDER BY id
        FOR UPDATE
    `

	rows, err := tx.Query(
		ctx,
		query,
		transfer.FromAccount,
		transfer.ToAccount,
	)
	if err != nil {
		return fmt.Errorf("unable to lock accounts: %w", err)
	}

	accounts := make(map[string]model.Account)

	for rows.Next() {
		var account model.Account

		if err := rows.Scan(
			&account.ID,
			&account.AccountNumber,
			&account.CustomerName,
			&account.Balance,
		); err != nil {
			rows.Close()
			return fmt.Errorf("unable to read account: %w", err)
		}

		accounts[account.AccountNumber] = account
	}

	if err := rows.Err(); err != nil {
		rows.Close()
		return fmt.Errorf("error reading accounts: %w", err)
	}

	rows.Close()

	fromAccount, ok := accounts[transfer.FromAccount]
	if !ok {
		return fmt.Errorf(
			"source account %s not found",
			transfer.FromAccount,
		)
	}

	toAccount, ok := accounts[transfer.ToAccount]
	if !ok {
		return fmt.Errorf(
			"destination account %s not found",
			transfer.ToAccount,
		)
	}

	if fromAccount.Balance < transfer.Amount {
		return errors.New("insufficient balance")
	}

	newFromBalance := fromAccount.Balance - transfer.Amount
	newToBalance := toAccount.Balance + transfer.Amount

	updateQuery := `
        UPDATE accounts
        SET balance = $1
        WHERE id = $2
    `

	if _, err := tx.Exec(
		ctx,
		updateQuery,
		newFromBalance,
		fromAccount.ID,
	); err != nil {
		return fmt.Errorf("unable to debit source account: %w", err)
	}

	if _, err := tx.Exec(
		ctx,
		updateQuery,
		newToBalance,
		toAccount.ID,
	); err != nil {
		return fmt.Errorf("unable to credit destination account: %w", err)
	}

	transactionQuery := `
        INSERT INTO transactions
        (from_account_id, to_account_id, transaction_type, amount)
        VALUES ($1, $2, $3, $4)
    `

	if _, err := tx.Exec(
		ctx,
		transactionQuery,
		fromAccount.ID,
		toAccount.ID,
		"MONEY_TRANSFER",
		transfer.Amount,
	); err != nil {
		return fmt.Errorf("unable to save transaction history: %w", err)
	}

	// Commit makes all balance and transaction changes permanent.
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("unable to commit transaction: %w", err)
	}

	return nil
}
