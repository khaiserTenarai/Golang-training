package service

import "context"

type AccountService interface {
	CreateAccount(ctx context.Context, name string, balance float64) error
	ListAccounts(ctx context.Context) error
	Transfer(ctx context.Context, fromID, toID int64, amount float64) error
}
