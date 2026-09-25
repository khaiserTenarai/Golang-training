package main

import "fmt"

func main() {
	var name string = "Abirami"
	var initial byte = 'S'
	var age int = 22
	var salary float32 = 70000.0
	var isWorking bool = true
	var startingLetter rune = 'A'

	fmt.Println("Go primitive data types")
	//String
	fmt.Println("Name: ", name)
	//byte
	fmt.Println("Initial: ", initial)
	//Integer
	fmt.Println("Age: ", age)
	//Float
	fmt.Println("Salary: ", salary)
	//Boolean
	fmt.Println("Working status: ", isWorking)
	//Rune
	fmt.Println("Starting Letter of the name: ", startingLetter)
}
