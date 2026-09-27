package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address3 struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode string `json:"pincode"`
}

type Department3 struct {
	Name    string `json:"name"`
	Manager string `json:"manager"`
}

type Employee5 struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Salary     float64    `json:"salary"`
	Address    Address3    `json:"address"`
	Department Department3 `json:"department"`
}

func main() {

	fmt.Println("\n********************************")
	fmt.Println("5. Unmarshal JSON into Employee.")
	fmt.Println("********************************")

	input := []byte(`{"id":101,"name":"Vimal Aditya","salary":500000,"address":{"city":"Bangalore","state":"Karnataka","pincode":"560016"},"department":{"name":"Engineering","manager":"Alex"}}`)
	
	var json_employee Employee5

	if err := json.Unmarshal(input, &json_employee); err != nil{
		log.Fatal(err)
	}

	fmt.Println(json_employee.ID)
	fmt.Println(json_employee.Name)

}