package main

import "fmt"

func main() {

	fmt.Println("\n*********************************************************")
	fmt.Println("6. Create a program using comparison and logical operators.")
	fmt.Println("***********************************************************")

	var age int
	var experience int
	var hasDegree bool

	fmt.Print("Enter Age: ")
	fmt.Scan(&age)

	fmt.Print("Enter Years of Experience: ")
	fmt.Scan(&experience)

	fmt.Print("Has College Degree? (true/false): ")
	fmt.Scan(&hasDegree)

	isAdult := age >= 18
	hasEnoughExp := experience >= 2

	fmt.Println("\n--- Evaluation Checks ---")
	fmt.Println("Is adult (Age >= 18):", isAdult)
	fmt.Println("Has enough experience (Exp >= 2):", hasEnoughExp)

	if isAdult && (hasEnoughExp || hasDegree) {
		fmt.Println("\nStatus: Candidate is Eligible for Hire")
	}

	if !isAdult || (experience == 0 && !hasDegree) {
		fmt.Println("Status: Candidate Needs Further Training")
	}
}