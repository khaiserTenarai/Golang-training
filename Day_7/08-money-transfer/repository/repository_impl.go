package repository
import (
	"context"
	"errors"

	"money_transfer/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransferRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewTransferRepository(
	db *pgxpool.Pool,
) TransferRepository {

	return &TransferRepositoryImpl{
		db: db,
	}
}

func (r *TransferRepositoryImpl) TransferMoney(
	transfer model.Transfer,
) error {

	ctx := context.Background()

	// Start transaction
	tx, err := r.db.Begin(ctx)

	if err != nil {
		return err
	}

	// Rollback automatically if something fails.
	defer tx.Rollback(ctx)

	// Check sender and lock the row.
	var senderBalance float64

	query := `
		SELECT balance
		FROM accounts
		WHERE id = $1
		FOR UPDATE
	`

	err = tx.QueryRow(
		ctx,
		query,
		transfer.FromAccountID,
	).Scan(&senderBalance)

	if err != nil {
		return errors.New("sender account not found")
	}

	// Check receiver.
	var receiverID int

	query = `
		SELECT id
		FROM accounts
		WHERE id = $1
		FOR UPDATE
	`

	err = tx.QueryRow(
		ctx,
		query,
		transfer.ToAccountID,
	).Scan(&receiverID)

	if err != nil {
		return errors.New("receiver account not found")
	}

	// Check balance.
	if senderBalance < transfer.Amount {
		return errors.New("insufficient balance")
	}

	// Deduct money from sender.
	query = `
		UPDATE accounts
		SET balance = balance - $1
		WHERE id = $2
	`

	_, err = tx.Exec(
		ctx,
		query,
		transfer.Amount,
		transfer.FromAccountID,
	)

	if err != nil {
		return err
	}

	// Add money to receiver.
	query = `
		UPDATE accounts
		SET balance = balance + $1
		WHERE id = $2
	`

	_, err = tx.Exec(
		ctx,
		query,
		transfer.Amount,
		transfer.ToAccountID,
	)

	if err != nil {
		return err
	}

	// Everything worked.
	err = tx.Commit(ctx)

	if err != nil {
		return err
	}

	return nil
}
