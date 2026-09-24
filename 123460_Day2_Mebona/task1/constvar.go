package main

import (
	"fmt"
)

const Category = "Vehicle"

const (
	Sedan     = iota 
	SUV              
	Hatchback       
)

func main() {
	
	var brand string = "Toyota"
	var year int = 2023

	var model, color string = "Camry", "Silver"

	price := 28500.99

	year = 2024 

	fmt.Println("Vehicle Details:")
	fmt.Println("Brand:", brand)
	fmt.Println("Model & Color:", model, ",", color)
	fmt.Println("Year:", year)
	fmt.Println("Price: $", price)

	fmt.Println("Category:", Category)
	fmt.Println("Vehicle Types (Sedan, SUV, Hatchback):", Sedan, SUV, Hatchback)
}