package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task07_bank_account_system/models"
)

type AccountView struct {
	scanner *bufio.Scanner
}

func NewAccountView() *AccountView {
	return &AccountView{scanner: bufio.NewScanner(os.Stdin)}
}

func (v *AccountView) readLine() string {
	v.scanner.Scan()
	return strings.TrimSpace(v.scanner.Text())
}

func (v *AccountView) ShowMenu() int {
	fmt.Println("\n===== BANK ACCOUNT SYSTEM =====")
	fmt.Println("1. Create Account")
	fmt.Println("2. List Accounts")
	fmt.Println("3. Check Balance")
	fmt.Println("4. Deposit")
	fmt.Println("5. Withdraw")
	fmt.Println("6. Transaction History")
	fmt.Println("7. Exit")
	fmt.Print("Enter choice: ")
	choice, _ := strconv.Atoi(v.readLine())
	return choice
}

func (v *AccountView) GetAccountInput() models.Account {
	var acc models.Account
	fmt.Print("Enter Holder Name: ")
	acc.HolderName = v.readLine()
	fmt.Print("Enter Account Type (savings/current): ")
	acc.AccountType = v.readLine()
	fmt.Print("Enter Initial Deposit: ")
	acc.Balance, _ = strconv.ParseFloat(v.readLine(), 64)
	return acc
}

func (v *AccountView) GetID() int {
	fmt.Print("Enter Account ID: ")
	id, _ := strconv.Atoi(v.readLine())
	return id
}

func (v *AccountView) GetAmount() float64 {
	fmt.Print("Enter Amount: ")
	amt, _ := strconv.ParseFloat(v.readLine(), 64)
	return amt
}

func (v *AccountView) GetDescription() string {
	fmt.Print("Enter Description: ")
	return v.readLine()
}

func (v *AccountView) ShowAccount(acc models.Account) {
	fmt.Printf("\nID: %d | Holder: %s | Type: %s | Balance: %.2f | Created: %s\n",
		acc.ID, acc.HolderName, acc.AccountType, acc.Balance, acc.CreatedAt.Format("2006-01-02 15:04"))
}

func (v *AccountView) ShowAccounts(accounts []models.Account) {
	if len(accounts) == 0 {
		fmt.Println("\nNo accounts found.")
		return
	}
	fmt.Printf("\n%-5s %-20s %-12s %-15s\n", "ID", "Holder", "Type", "Balance")
	fmt.Println(strings.Repeat("-", 57))
	for _, a := range accounts {
		fmt.Printf("%-5d %-20s %-12s %-15.2f\n", a.ID, a.HolderName, a.AccountType, a.Balance)
	}
}

func (v *AccountView) ShowBalance(acc models.Account) {
	fmt.Printf("\nAccount: %s | Balance: %.2f\n", acc.HolderName, acc.Balance)
}

func (v *AccountView) ShowTransactions(txns []models.Transaction) {
	if len(txns) == 0 {
		fmt.Println("\nNo transactions found.")
		return
	}
	fmt.Printf("\n%-5s %-12s %-12s %-15s %-25s %-20s\n", "ID", "Type", "Amount", "Balance After", "Description", "Date")
	fmt.Println(strings.Repeat("-", 95))
	for _, t := range txns {
		fmt.Printf("%-5d %-12s %-12.2f %-15.2f %-25s %-20s\n",
			t.ID, t.TxnType, t.Amount, t.BalanceAfter, t.Description, t.CreatedAt.Format("2006-01-02 15:04"))
	}
}

func (v *AccountView) ShowSuccess(msg string) {
	fmt.Println("\n[SUCCESS]", msg)
}

func (v *AccountView) ShowError(context string, err error) {
	fmt.Printf("\n[ERROR] %s: %v\n", context, err)
}
