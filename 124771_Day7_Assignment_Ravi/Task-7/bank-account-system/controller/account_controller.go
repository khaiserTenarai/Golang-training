package controller

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "bank-account-system/model"
    "bank-account-system/service"
)

type AccountController struct {
    accountService service.AccountService
    reader         *bufio.Reader
}

func NewAccountController(
    accountService service.AccountService,
) *AccountController {
    return &AccountController{
        accountService: accountService,
        reader:         bufio.NewReader(os.Stdin),
    }
}

func (c *AccountController) Start() {
    for {
        fmt.Println()
        fmt.Println("================================")
        fmt.Println("       BANK ACCOUNT SYSTEM")
        fmt.Println("================================")
        fmt.Println("1. Create Account")
        fmt.Println("2. Deposit")
        fmt.Println("3. Withdraw")
        fmt.Println("4. Balance Enquiry")
        fmt.Println("5. Transaction History")
        fmt.Println("6. Exit")
        fmt.Println("================================")

        choice := c.readInt("Enter choice: ")

        switch choice {
        case 1:
            c.createAccount()
        case 2:
            c.deposit()
        case 3:
            c.withdraw()
        case 4:
            c.balanceEnquiry()
        case 5:
            c.transactionHistory()
        case 6:
            fmt.Println("Thank you for using Bank Account System.")
            return
        default:
            fmt.Println("Invalid choice.")
        }
    }
}

func (c *AccountController) createAccount() {
    account := model.Account{
        AccountNumber: c.readString("Enter account number: "),
        CustomerName:  c.readString("Enter customer name: "),
        Balance:       c.readFloat("Enter initial balance: "),
    }

    err := c.accountService.CreateAccount(account)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Account created successfully.")
}

func (c *AccountController) deposit() {
    accountNumber := c.readString("Enter account number: ")
    amount := c.readFloat("Enter deposit amount: ")

    err := c.accountService.Deposit(accountNumber, amount)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Deposit successful.")
}

func (c *AccountController) withdraw() {
    accountNumber := c.readString("Enter account number: ")
    amount := c.readFloat("Enter withdrawal amount: ")

    err := c.accountService.Withdraw(accountNumber, amount)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Withdrawal successful.")
}

func (c *AccountController) balanceEnquiry() {
    accountNumber := c.readString("Enter account number: ")

    balance, err := c.accountService.GetBalance(accountNumber)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Current Balance: ₹%.2f
", balance)
}

func (c *AccountController) transactionHistory() {
    accountNumber := c.readString("Enter account number: ")

    transactions, err :=
        c.accountService.GetTransactionHistory(accountNumber)

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(transactions) == 0 {
        fmt.Println("No transactions found.")
        return
    }

    fmt.Println()
    fmt.Println("========================================")
    fmt.Println("          TRANSACTION HISTORY")
    fmt.Println("========================================")

    for _, transaction := range transactions {
        fmt.Printf(
            "%-10s ₹%-10.2f %s
",
            transaction.TransactionType,
            transaction.Amount,
            transaction.CreatedAt.Format("2006-01-02 15:04:05"),
        )
    }
}

func (c *AccountController) readString(prompt string) string {
    fmt.Print(prompt)
    value, _ := c.reader.ReadString('
')
    return strings.TrimSpace(value)
}

func (c *AccountController) readInt(prompt string) int {
    for {
        value := c.readString(prompt)
        number, err := strconv.Atoi(value)

        if err == nil {
            return number
        }

        fmt.Println("Please enter a valid number.")
    }
}

func (c *AccountController) readFloat(prompt string) float64 {
    for {
        value := c.readString(prompt)
        number, err := strconv.ParseFloat(value, 64)

        if err == nil {
            return number
        }

        fmt.Println("Please enter a valid amount.")
    }
}
