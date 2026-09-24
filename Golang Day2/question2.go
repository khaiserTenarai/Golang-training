package main

import "fmt"

func main() {
	// 1. STRING (Text)
	// Strings must always be wrapped in double quotes. 
	var heroName string = "Batman"

	// 2. INTEGER (Whole Numbers)
	// 'int' is used for counting things. No decimal points allowed
	var level int = 42

	// 3. FLOAT (Decimal Numbers)
	// 'float64' is used for precision (money, measurements, percentages)
	var healthPoints float64 = 98.5

	// 4. BOOLEAN (True/False)
	// 'bool' can only ever be true or false
	var isInvisible bool = true

	fmt.Println("Hero Name: ", heroName)
	fmt.Println("Current Level: ", level)
	fmt.Println("Health Points:", healthPoints)
	fmt.Println("Is Invisible: ", isInvisible)
}