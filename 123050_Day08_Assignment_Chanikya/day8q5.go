package main

import "fmt"

func main() {
	employee := struct {
		Name   string
		Salary float64
	}{
		Name:   "ram",
		Salary: 50000,
	}
	fmt.Println(employee.Name, employee.Salary)
}

//gofmt automatically formats your Go code into the standard Go style.
//PS C:\Users\M.Chanikya\Downloads\Overture_Project\data\GO\Assignment1\123050_Day08_Assignment_Chanikya> gofmt -w .\day8q5.go
