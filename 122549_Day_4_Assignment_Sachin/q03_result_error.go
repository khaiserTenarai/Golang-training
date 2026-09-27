// 3. Create functions returning (result, error).

package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID   int
	Name string
}

var employees = map[int]Employee{
	101: {ID: 101, Name: "John"},
	102: {ID: 102, Name: "Jane"},
}

func findEmployee(id int) (Employee, error) {
	emp, ok := employees[id]
	if !ok {
		return Employee{}, errors.New("employee not found")
	}
	return emp, nil
}

func main() {
	emp, err := findEmployee(101)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Found:", emp.Name)
	}

	_, err = findEmployee(999)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
