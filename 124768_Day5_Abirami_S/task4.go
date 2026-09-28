package main

import (
	"encoding/json"
	"fmt"
)

type Employee4 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	employee := Employee4{
		ID:   101,
		Name: "Abirami",
		Age:  22,
	}
	data, err := json.Marshal(employee)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
