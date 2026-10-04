package controller

import (
	"fmt"
	"task08_money_transfer/repository"
	"task08_money_transfer/view"
)

type TransferController struct {
	repo *repository.TransferRepository
	view *view.TransferView
}

func NewTransferController(repo *repository.TransferRepository, v *view.TransferView) *TransferController {
	return &TransferController{repo: repo, view: v}
}

func (c *TransferController) Run() {
	for {
		choice := c.view.ShowMenu()
		switch choice {
		case 1:
			c.CreateAccount()
		case 2:
			c.ListAccounts()
		case 3:
			c.TransferMoney()
		case 4:
			c.ViewTransferLogs()
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *TransferController) CreateAccount() {
	acc := c.view.GetAccountInput()
	id, err := c.repo.CreateAccount(acc)
	if err != nil {
		c.view.ShowError("creating account", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Account created with ID: %d", id))
}

func (c *TransferController) ListAccounts() {
	accounts, err := c.repo.GetAll()
	if err != nil {
		c.view.ShowError("listing accounts", err)
		return
	}
	c.view.ShowAccounts(accounts)
}

func (c *TransferController) TransferMoney() {
	fromID, toID, amount, desc := c.view.GetTransferInput()
	if err := c.repo.Transfer(fromID, toID, amount, desc); err != nil {
		c.view.ShowError("transfer (rolled back)", err)
		return
	}
	c.view.ShowSuccess(fmt.Sprintf("Transferred %.2f from Account #%d to Account #%d", amount, fromID, toID))
}

func (c *TransferController) ViewTransferLogs() {
	logs, err := c.repo.GetTransferLogs()
	if err != nil {
		c.view.ShowError("fetching transfer logs", err)
		return
	}
	c.view.ShowTransferLogs(logs)
}
