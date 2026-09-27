package main

import (
	"encoding/json"
	"fmt"
)
type Employee struct{
	ID int           `json:"id"`
	Name string      `json:"name"`
	Email string	 `json:"email"`
	Age int			 `json:"age"`
	Salary float64   `json:"salary"`
}

func main() {
	employee := Employee{
		ID : 101,
		Name : "Piyush",
		Email : "piyush@gmail.com",
		Age : 21,
		Salary:  25000.50,
	}
	 // Convert Employee object into JSON
	data, err := json.Marshal(employee)

	if err != nil{
		fmt.Println(err)
		return 
	}
	 // Convert []Byte into string and print
	input := []byte(`{"id": 102, "name":"Prachi", "email":"prachi@gmail.com", "age": 20, "Salary":100000}`)
	 
	fmt.Println(string(data))
	if err := json.Unmarshal(input, &employee); err != nil{
		fmt.Println(err)
	}
	fmt.Println("NAME:",employee.Name)
}
