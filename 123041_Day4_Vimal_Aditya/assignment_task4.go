package main

import (
	"errors"
	"fmt"
)

func calculateSalary(salaries ...float64) (float64, error){

	if len(salaries) <= 0{
		return 0, errors.New("Salary is not provided")
	}

	total := 0.0

	for _, salary := range salaries{
		total += salary
	}

	return total, nil

}

func main(){

	fmt.Println("\n********************************************")
	fmt.Println("4. Create a variadic salary calculation function.")
	fmt.Println("**********************************************")

	var sal1, sal2, sal3 float64

	fmt.Print("Enter Salary 1: ")
	fmt.Scan(&sal1)

	fmt.Print("Enter Salary 2: ")
	fmt.Scan(&sal2)

	fmt.Print("Enter Salary 3: ")
	fmt.Scan(&sal3)

	total, err := calculateSalary(sal1, sal2, sal3)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Total Salary:", total)
	}

}
