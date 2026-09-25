package main

import (
	"fmt"
	"strconv"
)

const AppName = "Go Task"
const AppStatus = true

func main(){
	var name string= "Indu"
	var age int= 23
	company := "Tenarai"
	var salary float64 = 50000.50

	// Integer
	var smallNumber int8 = 100
	var largeNumber int64 = 100000

	// Unsigned Integer
	var count uint = 50

	// Floating-point
	var price float32 = 99.99

	// Boolean
	var isEmployee bool = true

	// Complex
	var complexNumber complex64 = 3 + 4i

	// Rune
	var letter rune = 'A'

	var integer int
	var decimal float64
	var text string
	var flag bool
	var character rune
	var number complex64

	fmt.Println("********************************")
	fmt.Println("Task 4")
	fmt.Println("********************************")
	
	str := "100"

	// String → Integer
	integer, err := strconv.Atoi(str)
	if err!=nil{
		fmt.Println("Entered input cannot be converted to interger", str)
	}
	fmt.Println("Converted integer value is:", integer)
	// Integer → Float
	floatValue := float64(integer)

	fmt.Println("********************************")
	fmt.Println("Task 1")
	fmt.Println("********************************")

	fmt.Println("Appname", AppName)
	fmt.Println("AppStatus", AppStatus)
	fmt.Println("name", name)
	fmt.Println("age", age)
	fmt.Println("company", company)
	fmt.Println("salary", salary)

	fmt.Println("********************************")
	fmt.Println("Task 2")
	fmt.Println("********************************")

	fmt.Println("int:", age)
	fmt.Println("int8:", smallNumber)
	fmt.Println("int64:", largeNumber)
	fmt.Println("uint:", count)
	fmt.Println("float32:", price)
	fmt.Println("float64:", salary)
	fmt.Println("string:", name)
	fmt.Println("bool:", isEmployee)
	fmt.Println("complex64:", complexNumber)
	fmt.Println("rune:", letter)

	fmt.Println("********************************")
	fmt.Println("Task 3")
	fmt.Println("********************************")

	fmt.Println("int:", integer)
	fmt.Println("float64:", decimal)
	fmt.Println("string:", text)
	fmt.Println("bool:", flag)
	fmt.Println("rune:", character)
	fmt.Println("complex64:", number)

	fmt.Println("********************************")
	fmt.Println("Task 4")
	fmt.Println("********************************")

	fmt.Println("String:", str)
	fmt.Println("Integer:", integer)
	fmt.Println("Float:", floatValue)
}