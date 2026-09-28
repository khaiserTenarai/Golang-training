package main

import "fmt"

type Employee3 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	employee := Employee3{
		ID:   101,
		Name: "Abirami",
		Age:  22,
	}
	fmt.Printf("%+v\n", employee)
}
