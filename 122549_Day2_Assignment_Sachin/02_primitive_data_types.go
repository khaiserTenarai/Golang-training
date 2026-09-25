// Day 2 - Task 2: Primitive Data Types
package main

import "fmt"

func main() {
	var isActive bool = true
	var name string = "Sneha"
	var age int = 25
	var smallNumber int8 = 100
	var year int16 = 2026
	var population int32 = 140000000
	var distanceKM int64 = 384400
	var count uint = 10
	var price float32 = 499.99
	var pi float64 = 3.14159265
	var initial byte = 'S'
	var letter rune = 'A'

	fmt.Println("bool    :", isActive)
	fmt.Println("string  :", name)
	fmt.Println("int     :", age)
	fmt.Println("int8    :", smallNumber)
	fmt.Println("int16   :", year)
	fmt.Println("int32   :", population)
	fmt.Println("int64   :", distanceKM)
	fmt.Println("uint    :", count)
	fmt.Println("float32 :", price)
	fmt.Println("float64 :", pi)
	fmt.Println("byte    :", initial)
	fmt.Println("rune    :", letter)
}
