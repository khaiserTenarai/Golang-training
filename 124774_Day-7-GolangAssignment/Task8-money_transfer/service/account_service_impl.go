package service

import (
	"money-transfer/repository"
	"money-transfer/utility"
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

func (s *AccountServiceImpl) Transfer(
	fromID int,
	toID int,
	amount float64,
) error {

	// Validation

	if err := utility.ValidateAccountID(fromID); err != nil {
		return err
	}

	if err := utility.ValidateAccountID(toID); err != nil {
		return err
	}

	if err := utility.ValidateAmount(amount); err != nil {
		return err
	}

	if err := utility.ValidateDifferentAccounts(
		fromID,
		toID,
	); err != nil {
		return err
	}

	// Call repository

	return s.repository.Transfer(
		fromID,
		toID,
		amount,
	)
}
