package service

import (
	"errors"

	"bankaccount/model"
	"bankaccount/repository"
	"bankaccount/utility"
)

type AccountServiceImpl struct {
	repository repository.AccountRepository
}

func NewAccountService(
	repository repository.AccountRepository,
) AccountService {

	return &AccountServiceImpl{
		repository: repository,
	}
}

// CREATE ACCOUNT

func (s *AccountServiceImpl) CreateAccount(
	account model.Account,
) error {

	// Validation

	if err := utility.ValidateName(account.Name); err != nil {
		return err
	}

	if err := utility.ValidateOpeningBalance(
		account.Balance,
	); err != nil {
		return err
	}

	// Repository

	return s.repository.Create(account)
}

// DEPOSIT

func (s *AccountServiceImpl) Deposit(
	id int,
	amount float64,
) error {

	if err := utility.ValidateAccountID(id); err != nil {
		return err
	}

	if err := utility.ValidateAmount(amount); err != nil {
		return err
	}

	return s.repository.Deposit(id, amount)
}

// WITHDRAW

func (s *AccountServiceImpl) Withdraw(
	id int,
	amount float64,
) error {

	if err := utility.ValidateAccountID(id); err != nil {
		return err
	}

	if err := utility.ValidateAmount(amount); err != nil {
		return err
	}

	return s.repository.Withdraw(id, amount)
}

// BALANCE ENQUIRY

func (s *AccountServiceImpl) FindAccount(
	id int,
) (model.Account, error) {

	if id <= 0 {
		return model.Account{},
			errors.New("account ID must be greater than 0")
	}

	return s.repository.FindByID(id)
}

// TRANSACTION HISTORY

func (s *AccountServiceImpl) FindTransactions(
	id int,
) ([]model.Transaction, error) {

	if err := utility.ValidateAccountID(id); err != nil {
		return nil, err
	}

	return s.repository.FindTransactions(id)
}
