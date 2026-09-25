package main

import (
	"errors"
	"fmt"
)

func main(){

	fmt.Println("\n**************************************************")
	fmt.Println("5. Create an anonymous function for salary filtering")
	fmt.Println("****************************************************")

	var minSalary int

	fmt.Print("Enter minimum salary to filter: ")
	fmt.Scan(&minSalary)

	filterSalary := func (minSal int, salaries ...int)  ([]int, error){
		if minSal <= 0{
			return nil, errors.New("Minimum salary cannot be less than or equal to 0 ")
		} 

		var filtered []int

		for _, salary := range salaries{
			if salary >= minSalary{
				filtered = append(filtered, salary)
			}
		}

		if len(filtered) == 0 {
			return nil, errors.New("No salaries met the minimum filter")
		}

		return filtered, nil
	}

	result, err := filterSalary(minSalary, 1000, 2000, 3000, 4000, 5000, 6000)
	
	if err != nil{
		fmt.Println("Error:", err)
	}else{
		fmt.Println("Filtered:", result)
	}

}
