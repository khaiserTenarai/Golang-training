package service

import (
	"bank-account-system/model"
	"bank-account-system/repository"
)

type AccountServiceImpl struct {
	repository repository.AccountRepository
}

func NewAccountService(repository repository.AccountRepository) AccountService {
	return &AccountServiceImpl{
		repository: repository,
	}
}

func (s *AccountServiceImpl) CreateAccount(account model.Account) error {
	return s.repository.CreateAccount(account)
}

func (s *AccountServiceImpl) GetAccount(id int) (*model.Account, error) {
	return s.repository.GetAccount(id)
}

func (s *AccountServiceImpl) Deposit(id int, amount float64) error {
	return s.repository.Deposit(id, amount)
}

func (s *AccountServiceImpl) Withdraw(id int, amount float64) error {
	return s.repository.Withdraw(id, amount)
}

func (s *AccountServiceImpl) GetTransactionHistory(accountID int) ([]model.Transaction, error) {
	return s.repository.GetTransactionHistory(accountID)
}
