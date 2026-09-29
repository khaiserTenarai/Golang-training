package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
	Salary float64 `json:"salary"`
}

func main() {
	employee := Employee{
		ID: 101,
		Name: "Indu",
		Email: "indu@gmail.com",
		Age: 23,
		Salary: 50000,
	}

	data, err:=json.Marshal(employee)

	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println(string(data))

}
	





