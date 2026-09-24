package main

import (
	"fmt"
)

func main() {
	
	var age int
	var isCitizen bool
	var isRegistered bool

	fmt.Println("--- Voting Eligibility Checker ---")

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Print("Are you a citizen? (true or false): ")
	fmt.Scan(&isCitizen)

	fmt.Print("Are you registered to vote? (true or false): ")
	fmt.Scan(&isRegistered)


	isOldEnough := age >= 18
	fmt.Println("Old enough (18+)?", isOldEnough)

	
	isEligible := isOldEnough && isCitizen
	fmt.Println("General Eligibility (Age AND Citizen)?", isEligible)

	
	canVoteToday := isEligible && isRegistered
	fmt.Println("Can cast a ballot today (Eligible AND Registered)?", canVoteToday)


	if !canVoteToday {
		fmt.Println("\nStatus: You are NOT ready to vote today. Please resolve your age, citizenship, or registration.")
	} else {
		fmt.Println("\nStatus: You are all set to vote!")
	}
}