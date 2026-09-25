package main

import "fmt"

func main() {
	age, exp := 22, 1

	fmt.Println("Is Adult?:", age >= 18)
	fmt.Println("Has Exp:", exp > 0)
	fmt.Println("Eligible:", age >= 18 && exp > 0)
}