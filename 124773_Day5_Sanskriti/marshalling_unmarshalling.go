package main

import (
	"encoding/json"
	"fmt"
	"log"
	//"strings"
)

type Emp struct{
	ID int `json:"id`
	Name string  `json:"Name"`
	Email string `json: "email"`
	Age int  `json:"age"`
	Salary float64 `json: "salary"`
}

func main(){
	emp := Emp{
		ID: 1,
		Name: "Sans",
		Email: "sans@gmail.com",
		Age: 22,
		Salary: 33000.0,
	}

	data,err:= json.Marshal(emp)
	if err != nil{
		fmt.Println(err)
		return
}	

	fmt.Println(string(data))

	//input := []data({})
	if err := json.Unmarshal(data, &emp); err != nil {
		log.Fatal(err)
	}
	fmt.Println(emp.Name)

}