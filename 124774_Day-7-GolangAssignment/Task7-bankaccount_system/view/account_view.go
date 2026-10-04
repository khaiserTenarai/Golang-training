package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"bankaccount/model"
)

type AccountView struct {
	reader *bufio.Reader
}

func NewAccountView() *AccountView {

	return &AccountView{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (v *AccountView) ShowMenu() {

	fmt.Println()
	fmt.Println("========== BANK ACCOUNT SYSTEM ==========")
	fmt.Println("1. Create Account")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Balance Enquiry")
	fmt.Println("5. Transaction History")
	fmt.Println("6. Exit")
	fmt.Println("=========================================")
}

func (v *AccountView) ReadInt(message string) int {

	for {

		fmt.Print(message)

		input, _ := v.reader.ReadString('\n')

		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)

		if err == nil {
			return value
		}

		fmt.Println("Please enter a valid number.")
	}
}

func (v *AccountView) ReadFloat(message string) float64 {

	for {

		fmt.Print(message)

		input, _ := v.reader.ReadString('\n')

		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)

		if err == nil {
			return value
		}

		fmt.Println("Please enter a valid amount.")
	}
}

func (v *AccountView) ReadString(message string) string {

	fmt.Print(message)

	input, _ := v.reader.ReadString('\n')

	return strings.TrimSpace(input)
}

// CREATE ACCOUNT

func (v *AccountView) ReadAccount() model.Account {

	return model.Account{
		Name:    v.ReadString("Enter Name: "),
		Balance: v.ReadFloat("Enter Opening Balance: "),
	}
}

// SHOW ACCOUNT

func (v *AccountView) ShowAccount(
	account model.Account,
) {

	fmt.Println("--------------------------------")
	fmt.Println("ID             :", account.ID)
	fmt.Println("Account Number :", account.AccountNumber)
	fmt.Println("Name           :", account.Name)
	fmt.Println("Balance        :", account.Balance)
	fmt.Println("--------------------------------")
}

// SHOW TRANSACTIONS

func (v *AccountView) ShowTransactions(
	transactions []model.Transaction,
) {

	if len(transactions) == 0 {
		fmt.Println("No transactions found.")
		return
	}

	fmt.Println("\n========== TRANSACTION HISTORY ==========")

	for _, transaction := range transactions {

		fmt.Println("--------------------------------")
		fmt.Println("Transaction ID :", transaction.ID)
		fmt.Println("Type           :", transaction.Type)
		fmt.Println("Amount         :", transaction.Amount)
		fmt.Println("Balance After  :", transaction.BalanceAfter)
	}
}

func (v *AccountView) ShowMessage(
	message string,
) {

	fmt.Println(message)
}
