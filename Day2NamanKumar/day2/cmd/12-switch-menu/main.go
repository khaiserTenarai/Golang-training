// Task 12: A menu-driven program using switch: a simple bank account.
// It shows every form of switch: value, multiple values, no-condition
// (tagless), init statement, fallthrough, and type switch.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var balance = 10000.0
var history []string

func main() {
	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(`
====== BANK MENU ======
 1. Check balance
 2. Deposit
 3. Withdraw
 4. Transaction history
 5. Account tier
 6. Switch examples
 0. Exit`)
		choice := read(in, "Choose: ")

		switch choice { // value switch: no "break" needed, Go stops automatically
		case "1":
			fmt.Printf("Balance: ₹%.2f\n", balance)
		case "2":
			deposit(readAmount(in))
		case "3":
			withdraw(readAmount(in))
		case "4":
			if len(history) == 0 {
				fmt.Println("No transactions yet.")
			}
			for i, h := range history {
				fmt.Printf("%d. %s\n", i+1, h)
			}
		case "5":
			fmt.Println("Tier:", tier(balance))
		case "6":
			examples()
		case "0", "q", "exit": // several values in one case
			fmt.Println("Thank you for banking with us!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func deposit(amt float64) {
	switch { // tagless switch: works like if / else-if
	case amt <= 0:
		fmt.Println("Amount must be positive.")
	case amt > 200000:
		fmt.Println("Deposits above ₹2,00,000 need branch approval.")
	default:
		balance += amt
		history = append(history, fmt.Sprintf("Deposit  +₹%.2f", amt))
		fmt.Printf("Deposited. New balance: ₹%.2f\n", balance)
	}
}

func withdraw(amt float64) {
	switch remaining := balance - amt; { // switch with an init statement
	case amt <= 0:
		fmt.Println("Amount must be positive.")
	case remaining < 0:
		fmt.Printf("Insufficient funds. Balance is ₹%.2f\n", balance)
	case remaining < 1000:
		fmt.Println("Denied: minimum balance of ₹1,000 must be kept.")
	default:
		balance = remaining
		history = append(history, fmt.Sprintf("Withdraw -₹%.2f", amt))
		fmt.Printf("Withdrawn. New balance: ₹%.2f\n", balance)
	}
}

func tier(b float64) string {
	switch {
	case b >= 100000:
		return "Platinum"
	case b >= 50000:
		return "Gold"
	case b >= 10000:
		return "Silver"
	default:
		return "Basic"
	}
}

func examples() {
	fmt.Println("\n-- fallthrough: runs the NEXT case too (rarely used) --")
	fmt.Print("Benefits for Gold: ")
	switch "Gold" {
	case "Platinum":
		fmt.Print("lounge access, ")
		fallthrough
	case "Gold":
		fmt.Print("free cheque book, ")
		fallthrough
	case "Silver":
		fmt.Print("debit card")
	}
	fmt.Println()

	fmt.Println("\n-- type switch: branch on the dynamic type --")
	for _, v := range []any{42, 3.14, "hello", true, []int{1, 2}} {
		switch x := v.(type) {
		case int:
			fmt.Println(x, "is an int")
		case float64:
			fmt.Println(x, "is a float64")
		case string:
			fmt.Printf("%q is a string of length %d\n", x, len(x))
		case bool:
			fmt.Println(x, "is a bool")
		default:
			fmt.Printf("%v is %T\n", x, x)
		}
	}
}

func read(in *bufio.Scanner, label string) string {
	fmt.Print(label)
	if !in.Scan() {
		fmt.Println("\nGoodbye!")
		os.Exit(0)
	}
	return strings.TrimSpace(in.Text())
}

func readAmount(in *bufio.Scanner) float64 {
	v, err := strconv.ParseFloat(read(in, "Amount: ₹"), 64)
	if err != nil {
		fmt.Println("Not a valid number.")
		return 0
	}
	return v
}
