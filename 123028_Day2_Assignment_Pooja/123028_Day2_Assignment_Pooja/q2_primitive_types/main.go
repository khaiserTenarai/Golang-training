package main

import "fmt"

func main() {
	var name string = "Pooja"
	var MobileNumber int = 998877665544
	var smallnumber int8 = 100
	var largenumber int64 = 1000000
	var temperature float32 = 33.5
	var salary float64 = 75000.50
	var isemployee bool = true
	var data byte = 65
	var grade rune = 'A'

	fmt.Println("Name:", name)
	fmt.Println("Mobile Number:", MobileNumber)
	fmt.Println("Small number:", smallnumber)
	fmt.Println("Large number:", largenumber)
	fmt.Println("Temperature:", temperature)
	fmt.Println("Salary:", salary)
	fmt.Println("Is employee:", isemployee)
	fmt.Println("Byte value:", data)
	fmt.Println("Grade:", grade)
}
