package main

import "fmt"

func main() {

	var age int
	var exp int
	var degree bool

	fmt.Print("Enter age: ")
	fmt.Scan(&age)

	fmt.Print("Enter experience in years: ")
	fmt.Scan(&exp)

	fmt.Print("If having degree mention true or false: ")
	fmt.Scan(&degree)

	job := degree && exp >= 1
	fmt.Println("Eligible for job:", job)

	intern := degree || exp >= 1
	fmt.Println("Eligible for internship:", intern)

	fresher := !(exp > 0)
	fmt.Println("he is fresher:", fresher)
}
