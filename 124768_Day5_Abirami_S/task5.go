package main

import (
	"encoding/json"
	"fmt"
)

type Employee5 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	employeejson := `{
		"id":   101,
		"name": "Abirami",
		"age":  22
	}`
	var employee Employee5

	err := json.Unmarshal([]byte(employeejson), &employee)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(employee)
}
