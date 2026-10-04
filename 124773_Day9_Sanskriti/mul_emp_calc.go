package main

import (
	"fmt"
	"time"
)

func calculateSalary(name string, salary float64) {
	time.Sleep(time.Second)
	fmt.Println(name, "salary =", salary)
}

func main() {
	go calculateSalary("Sanskriti", 50000)
	go calculateSalary("Rahul", 60000)
	go calculateSalary("Priya", 55000)

	time.Sleep(2 * time.Second)

	fmt.Println("All calculations completed")
}