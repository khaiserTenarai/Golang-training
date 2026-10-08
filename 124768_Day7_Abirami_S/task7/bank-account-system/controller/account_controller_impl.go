package controller

import (
	"bank-account-system/model"
	"bank-account-system/service"
)

type AccountControllerImpl struct {
	service service.AccountService
}

func NewAccountController(service service.AccountService) AccountController {
	return &AccountControllerImpl{
		service: service,
	}
}

func (c *AccountControllerImpl) CreateAccount(account model.Account) error {
	return c.service.CreateAccount(account)
}

func (c *AccountControllerImpl) GetAccount(id int) (*model.Account, error) {
	return c.service.GetAccount(id)
}

func (c *AccountControllerImpl) Deposit(id int, amount float64) error {
	return c.service.Deposit(id, amount)
}

func (c *AccountControllerImpl) Withdraw(id int, amount float64) error {
	return c.service.Withdraw(id, amount)
}

func (c *AccountControllerImpl) GetTransactionHistory(accountID int) ([]model.Transaction, error) {
	return c.service.GetTransactionHistory(accountID)
}
