package main

import "fmt"

func main() {
	age := 28
	experienceYears := 4
	hasDegree := true


	isEligibleForSeniorRole := age >= 25 && experienceYears >= 5 && hasDegree
	isEligibleForMidRole := (age >= 21 && experienceYears >= 2) || hasDegree
	needsTraining := !hasDegree || experienceYears < 2

	fmt.Printf("Eligible for Senior Role: %t\n", isEligibleForSeniorRole)
	fmt.Printf("Eligible for Mid Role:    %t\n", isEligibleForMidRole)
	fmt.Printf("Needs Training:           %t\n", needsTraining)
}