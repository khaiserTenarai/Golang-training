package main

import "fmt"

func main() {
	
	fmt.Println("--- 1. Factorial ---")
	numFact := 5
	result := 1

	for i := 1; i <= numFact; i++ {
		result = result * i
	}
	fmt.Println("Factorial of", numFact, "is:", result)

	fmt.Println("\n--- 2. Fibonacci Sequence ---")
	fibCount := 7 
	a := 0
	b := 1

	fmt.Print("First ", fibCount, " numbers: ")
	for i := 0; i < fibCount; i++ {
		fmt.Print(a, " ")
		next := a + b 
		a = b         
		b = next     
	}
	fmt.Println()

	
	fmt.Println("\n--- 3. Prime Number ---")
	numPrime := 7
	isPrime := true

	
	for i := 2; i < numPrime; i++ {
		if numPrime%i == 0 { 
			isPrime = false
			break
		}
	}

	if isPrime && numPrime > 1 {
		fmt.Println(numPrime, "is a Prime Number ")
	} else {
		fmt.Println(numPrime, "is NOT a Prime Number ")
	}

	
	fmt.Println("\n--- 4. Reverse Number ---")
	numToReverse := 1234
	reversed := 0

	
	for numToReverse > 0 {
		lastDigit := numToReverse % 10            
		reversed = (reversed * 10) + lastDigit     
		numToReverse = numToReverse / 10           
	}
	fmt.Println("Original was 1234. Reversed is:", reversed)

	fmt.Println("\n--- 5. Palindrome ---")
	originalPal := 1221
	tempPal := originalPal 
	reversedPal := 0

	for tempPal > 0 {
		lastDigit := tempPal % 10
		reversedPal = (reversedPal * 10) + lastDigit
		tempPal = tempPal / 10
	}

	if originalPal == reversedPal {
		fmt.Println(originalPal, "reads the same backwards. It IS a palindrome! ")
	} else {
		fmt.Println(originalPal, "does NOT read the same backwards. ")
	}
}