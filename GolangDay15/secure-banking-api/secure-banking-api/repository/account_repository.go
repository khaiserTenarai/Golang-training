package repository

import (
	"context"
	"database/sql"
	"errors"

	"secure-banking-api/model"
)

type AccountRepository interface {
	Create(ctx context.Context, a *model.Account) error
	FindByID(ctx context.Context, id int64) (*model.Account, error)
	List(ctx context.Context, userID int64, limit, offset int) ([]model.Account, int, error)
	SetDocument(ctx context.Context, id int64, path string) error
}

type accountRepository struct{ db *sql.DB }

func NewAccountRepository(db *sql.DB) AccountRepository { return &accountRepository{db: db} }

func (r *accountRepository) Create(ctx context.Context, a *model.Account) error {
	return r.db.QueryRowContext(ctx,
		`INSERT INTO accounts (user_id, holder_name, balance)
		 VALUES ($1, $2, $3) RETURNING id, created_at`,
		a.UserID, a.HolderName, a.Balance,
	).Scan(&a.ID, &a.CreatedAt)
}

func (r *accountRepository) FindByID(ctx context.Context, id int64) (*model.Account, error) {
	var a model.Account
	err := r.db.QueryRowContext(ctx,
		`SELECT id, user_id, holder_name, balance, COALESCE(document_path, ''), created_at
		 FROM accounts WHERE id = $1`, id,
	).Scan(&a.ID, &a.UserID, &a.HolderName, &a.Balance, &a.DocumentPath, &a.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrNotFound
	}
	return &a, err
}

func (r *accountRepository) List(ctx context.Context, userID int64, limit, offset int) ([]model.Account, int, error) {
	var total int
	if err := r.db.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM accounts WHERE ($1 = 0 OR user_id = $1)`, userID,
	).Scan(&total); err != nil {
		return nil, 0, err
	}

	rows, err := r.db.QueryContext(ctx,
		`SELECT id, user_id, holder_name, balance, COALESCE(document_path, ''), created_at
		 FROM accounts WHERE ($1 = 0 OR user_id = $1)
		 ORDER BY id DESC LIMIT $2 OFFSET $3`, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var out []model.Account
	for rows.Next() {
		var a model.Account
		if err := rows.Scan(&a.ID, &a.UserID, &a.HolderName, &a.Balance, &a.DocumentPath, &a.CreatedAt); err != nil {
			return nil, 0, err
		}
		out = append(out, a)
	}
	return out, total, rows.Err()
}

func (r *accountRepository) SetDocument(ctx context.Context, id int64, path string) error {
	_, err := r.db.ExecContext(ctx, `UPDATE accounts SET document_path = $1 WHERE id = $2`, path, id)
	return err
}
