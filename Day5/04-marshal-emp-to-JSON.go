package main

import (
 "encoding/json"
 "fmt"
)

type Employee struct {
 ID int `json:"id"`
 Name string `json:"name"`
 Age int `json:"age"`
 Salary float64 `json:"salary"`
}

func main() {

 emp := Employee{
  ID: 101,
  Name: "Ganesh",
  Age: 22,
  Salary: 40000,
 }

 data, err := json.Marshal(emp)

 if err != nil {
  fmt.Println(err)
  return
 }

 fmt.Println(string(data))
}

/*
PS C:\Training\Go Lang\Day_5\124772_Day5(Go)_Assignment_Reddem_Ganesh_Reddy> go run .\04-marshal-emp-to-JSON.go
{"id":101,"name":"Ganesh","age":22,"salary":40000}
*/