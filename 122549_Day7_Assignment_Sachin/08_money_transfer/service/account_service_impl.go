package service

import (
	"context"
	"errors"
	"example.com/q8-money-transfer/repository"
)

type AccountServiceImpl struct{ repo repository.AccountRepository }

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &AccountServiceImpl{repo: repo}
}
func (s *AccountServiceImpl) CreateAccount(ctx context.Context, name string, balance float64) error {
	if name == "" || balance < 0 {
		return errors.New("invalid account details")
	}
	return s.repo.Create(ctx, name, balance)
}
func (s *AccountServiceImpl) ListAccounts(ctx context.Context) error { return s.repo.List(ctx) }
func (s *AccountServiceImpl) Transfer(ctx context.Context, fromID, toID int64, amount float64) error {
	if fromID == toID {
		return errors.New("source and destination must be different")
	}
	return s.repo.Transfer(ctx, fromID, toID, amount)
}
