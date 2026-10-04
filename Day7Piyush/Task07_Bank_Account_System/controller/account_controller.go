package controller

import (
	"fmt"
	"task07_bank_account_system/repository"
	"task07_bank_account_system/view"
)

type AccountController struct {
	repo *repository.AccountRepository
	view *view.AccountView
}

func NewAccountController(repo *repository.AccountRepository, v *view.AccountView) *AccountController {
	return &AccountController{repo: repo, view: v}
}

func (c *AccountController) Run() {
	for {
		choice := c.view.ShowMenu()
		switch choice {
		case 1:
			c.CreateAccount()
		case 2:
			c.ListAccounts()
		case 3:
			c.CheckBalance()
		case 4:
			c.Deposit()
		case 5:
			c.Withdraw()
		case 6:
			c.TransactionHistory()
		case 7:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *AccountController) CreateAccount() {
	acc := c.view.GetAccountInput()
	id, err := c.repo.CreateAccount(acc)
	if err != nil {
		c.view.ShowError("creating account", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Account created with ID: %d", id))
}

func (c *AccountController) ListAccounts() {
	accounts, err := c.repo.GetAll()
	if err != nil {
		c.view.ShowError("listing accounts", err)
		return
	}
	c.view.ShowAccounts(accounts)
}

func (c *AccountController) CheckBalance() {
	id := c.view.GetID()
	acc, err := c.repo.GetByID(id)
	if err != nil {
		c.view.ShowError("fetching account", err)
		return
	}
	c.view.ShowBalance(acc)
}

func (c *AccountController) Deposit() {
	id := c.view.GetID()
	amount := c.view.GetAmount()
	desc := c.view.GetDescription()
	if err := c.repo.Deposit(id, amount, desc); err != nil {
		c.view.ShowError("deposit", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Deposited %.2f successfully.", amount))
}

func (c *AccountController) Withdraw() {
	id := c.view.GetID()
	amount := c.view.GetAmount()
	desc := c.view.GetDescription()
	if err := c.repo.Withdraw(id, amount, desc); err != nil {
		c.view.ShowError("withdrawal", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Withdrawn %.2f successfully.", amount))
}

func (c *AccountController) TransactionHistory() {
	id := c.view.GetID()
	txns, err := c.repo.GetTransactionHistory(id)
	if err != nil {
		c.view.ShowError("fetching history", err)
		return
	}
	c.view.ShowTransactions(txns)
}
