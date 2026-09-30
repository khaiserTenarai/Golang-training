package repository
import (
	"context"
	"errors"

	"bank_account/model"

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

func (r *AccountRepositoryImpl) Create(
	account model.Account,
) error {

	query := `
		INSERT INTO accounts
		(name, email, balance)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		account.Name,
		account.Email,
		account.Balance,
	)

	return err
}

func (r *AccountRepositoryImpl) FindByID(
	id int,
) (model.Account, error) {

	var account model.Account

	query := `
		SELECT id, name, email, balance
		FROM accounts
		WHERE id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.Balance,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return account, errors.New("account not found")
	}

	return account, err
}

func (r *AccountRepositoryImpl) Deposit(
	id int,
	amount float64,
) error {

	tx, err := r.db.Begin(context.Background())

	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	query := `
		UPDATE accounts
		SET balance = balance + $1
		WHERE id = $2
	`

	result, err := tx.Exec(
		context.Background(),
		query,
		amount,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("account not found")
	}

	query = `
		INSERT INTO transactions
		(account_id, type, amount)
		VALUES ($1, $2, $3)
	`

	_, err = tx.Exec(
		context.Background(),
		query,
		id,
		"DEPOSIT",
		amount,
	)

	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

func (r *AccountRepositoryImpl) Withdraw(
	id int,
	amount float64,
) error {

	tx, err := r.db.Begin(context.Background())

	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	query := `
		UPDATE accounts
		SET balance = balance - $1
		WHERE id = $2
		  AND balance >= $1
	`

	result, err := tx.Exec(
		context.Background(),
		query,
		amount,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New(
			"insufficient balance or account not found",
		)
	}

	query = `
		INSERT INTO transactions
		(account_id, type, amount)
		VALUES ($1, $2, $3)
	`

	_, err = tx.Exec(
		context.Background(),
		query,
		id,
		"WITHDRAW",
		amount,
	)

	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

func (r *AccountRepositoryImpl) GetTransactions(
	id int,
) ([]model.Transaction, error) {

	query := `
		SELECT id, account_id, type, amount
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

	transactions := make([]model.Transaction, 0)

	for rows.Next() {

		var transaction model.Transaction

		err := rows.Scan(
			&transaction.ID,
			&transaction.AccountID,
			&transaction.Type,
			&transaction.Amount,
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
