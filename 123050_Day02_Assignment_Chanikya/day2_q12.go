package main

import "fmt"

func main() {
	var currentBalance float64 = 1000.00
	var userSelection int
	var runAgain string

	for {
		fmt.Println("\n************** AUTOMATED TELLER MACHINE **************")
		fmt.Println("1. Check Account Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Cash")
		fmt.Println("4. Reset Account Funds")
		fmt.Println("5. Exit System")

		fmt.Print("Select an option (1-5): ")
		fmt.Scan(&userSelection)

		switch userSelection {

		case 1:
			fmt.Printf("Your current balance is: ₹%.2f\n", currentBalance)

		case 2:
			var depositAmount float64
			fmt.Print("Enter amount to deposit: ₹")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Deposit must be greater than zero.")
			} else {
				currentBalance = currentBalance + depositAmount
				fmt.Printf("Successfully deposited. New Balance: ₹%.2f\n", currentBalance)
			}

		case 3:

			var withdrawalAmount float64
			fmt.Print("Enter amount to withdraw: ₹")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Withdrawal must be greater than zero.")
			} else if withdrawalAmount > currentBalance {
				fmt.Println("Insufficient funds! Transaction cancelled.")
			} else {
				currentBalance = currentBalance - withdrawalAmount
				fmt.Printf("Please collect your cash. Remaining Balance: ₹%.2f\n", currentBalance)
			}

		case 4:

			currentBalance = 0.00
			fmt.Println("Account balance has been completely cleared to zero.")

		case 5:

			fmt.Println("Thank you for using our banking services. Goodbye!")
			return

		default:
			fmt.Println("Invalid selection. Please choose a number between 1 and 5.")
		}

		fmt.Print("\nDo you want to perform another transaction? (yes/no): ")
		fmt.Scan(&runAgain)

		if runAgain != "yes" {
			fmt.Println("Thank you for using our banking services. Goodbye!")
			break
		}
	}
}
