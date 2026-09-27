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

 jsonData := `{"id":101,"name":"Ganesh","age":22,"salary":40000}`

 var emp Employee

 err := json.Unmarshal([]byte(jsonData), &emp)

 if err != nil {
  fmt.Println(err)
  return
 }

 fmt.Println(emp.ID)
 fmt.Println(emp.Name)
 fmt.Println(emp.Age)
 fmt.Println(emp.Salary)
}


/*
PS C:\Training\Go Lang\Day_5\124772_Day5(Go)_Assignment_Reddem_Ganesh_Reddy> go run .\05-unmarshal-JSON-to-emp.go
101
Ganesh
22
40000
*/