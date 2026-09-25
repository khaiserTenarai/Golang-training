// 6. Create a program using comparison and logical operators.

package main

import "fmt"

func main() {
	age := 24
	yearsOfExperience := 2

	isAdult := age >= 18
	isEligibleByAge := age >= 21
	hasEnoughExperience := yearsOfExperience >= 1

	fmt.Println("Is adult:", isAdult)
	fmt.Println("Age equal to 24:", age == 24)
	fmt.Println("Age not equal to 30:", age != 30)

	// logical AND
	canApply := isEligibleByAge && hasEnoughExperience
	fmt.Println("Eligible to apply for the job:", canApply)

	// logical OR
	needsReview := (age < 21) || (yearsOfExperience < 1)
	fmt.Println("Needs manual review:", needsReview)

	// logical NOT
	fmt.Println("Is NOT adult:", !isAdult)
}
