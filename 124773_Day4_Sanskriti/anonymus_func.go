package main

import "fmt"

func main() {
	salaries := []float64{2500, 45000, 60000, 30000, 75000}
	filter := func(salary float64) bool{
		return salary >4000
	}
	for _, salary := range salaries {
		if filter(salary){
			fmt.Println("salary" , salary)
		}
	}


}