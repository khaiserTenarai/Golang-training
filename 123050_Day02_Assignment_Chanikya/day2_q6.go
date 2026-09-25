package main

import "fmt"

func main() {
	var age int
	var isCitizen bool

	fmt.Println("Voting Eligibility Checker")
	fmt.Print("Enter your age: ")
	_, err := fmt.Scan(&age)
	if err != nil || age < 0 {
		fmt.Println("Invalid age entered.")
		return
	}

	fmt.Print("Are you a citizen? (true/false): ")
	_, err = fmt.Scan(&isCitizen)
	if err != nil {
		fmt.Println("Invalid input. Please enter true or false.")
		return
	}

	if age >= 18 && isCitizen {
		fmt.Println("Eligible to Vote")
	} else {
		fmt.Println("Not Eligible to Vote:")
	}

}
