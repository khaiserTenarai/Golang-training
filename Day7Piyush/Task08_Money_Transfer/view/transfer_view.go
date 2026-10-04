package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task08_money_transfer/models"
)

type TransferView struct {
	scanner *bufio.Scanner
}

func NewTransferView() *TransferView {
	return &TransferView{scanner: bufio.NewScanner(os.Stdin)}
}

func (v *TransferView) readLine() string {
	v.scanner.Scan()
	return strings.TrimSpace(v.scanner.Text())
}

func (v *TransferView) ShowMenu() int {
	fmt.Println("\n===== MONEY TRANSFER =====")
	fmt.Println("1. Create Account")
	fmt.Println("2. List Accounts")
	fmt.Println("3. Transfer Money")
	fmt.Println("4. View Transfer Logs")
	fmt.Println("5. Exit")
	fmt.Print("Enter choice: ")
	choice, _ := strconv.Atoi(v.readLine())
	return choice
}

func (v *TransferView) GetAccountInput() models.Account {
	var acc models.Account
	fmt.Print("Enter Holder Name: ")
	acc.HolderName = v.readLine()
	fmt.Print("Enter Initial Balance: ")
	acc.Balance, _ = strconv.ParseFloat(v.readLine(), 64)
	return acc
}

func (v *TransferView) GetTransferInput() (int, int, float64, string) {
	fmt.Print("From Account ID: ")
	fromID, _ := strconv.Atoi(v.readLine())
	fmt.Print("To Account ID: ")
	toID, _ := strconv.Atoi(v.readLine())
	fmt.Print("Amount: ")
	amount, _ := strconv.ParseFloat(v.readLine(), 64)
	fmt.Print("Description: ")
	desc := v.readLine()
	return fromID, toID, amount, desc
}

func (v *TransferView) ShowAccounts(accounts []models.Account) {
	if len(accounts) == 0 {
		fmt.Println("\nNo accounts found.")
		return
	}
	fmt.Printf("\n%-5s %-25s %-15s\n", "ID", "Holder", "Balance")
	fmt.Println(strings.Repeat("-", 48))
	for _, a := range accounts {
		fmt.Printf("%-5d %-25s %-15.2f\n", a.ID, a.HolderName, a.Balance)
	}
}

func (v *TransferView) ShowTransferLogs(logs []models.TransferLog) {
	if len(logs) == 0 {
		fmt.Println("\nNo transfer logs found.")
		return
	}
	fmt.Printf("\n%-4s %-15s %-15s %-12s %-10s %-20s %-20s\n", "ID", "From", "To", "Amount", "Status", "Description", "Date")
	fmt.Println(strings.Repeat("-", 100))
	for _, l := range logs {
		fmt.Printf("%-4d %-15s %-15s %-12.2f %-10s %-20s %-20s\n",
			l.ID, l.FromName, l.ToName, l.Amount, l.Status, l.Description, l.CreatedAt.Format("2006-01-02 15:04"))
	}
}

func (v *TransferView) ShowSuccess(msg string) {
	fmt.Println("\n[SUCCESS]", msg)
}

func (v *TransferView) ShowError(context string, err error) {
	fmt.Printf("\n[ERROR] %s: %v\n", context, err)
}
