package day4assignmentravi
package main

import "fmt"

func greet(name string, age int) {
	fmt.Println(name, age)
}

func calculate(a, b, c int) (int, int, int) {
	sum := a + b
	difference := a - b
	multiply := a * c
	return sum, difference, multiply
}

func calculate2(a, b int) (sum int, difference int) {
	sum = a + b
	difference = a - b
	return
}

func employee(name string, age int, salary float32) (string, int, float32) {
	return name, age, salary
}

func add(numbers ...int) int {
	total := 0
	// for _, number := range numbers {
	// 	total += number
	// }
	// return total
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

func main() {
	// greet("Ravi", 23)
	// a, b, c := calculate(10, 5, 2)
	// fmt.Println(a, b, c)
	// name, age, salary := employee("Ravi", 23, 50000)
	// fmt.Println(name, age, salary)
	// d, e := calculate2(3, 6)
	// fmt.Println(d, e)

	fmt.Println(add(1, 2, 3))

	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(2, 3))

	func() {
		fmt.Println("Hello")
	}()
}
