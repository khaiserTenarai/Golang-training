package main

import "fmt"

//Student  stores student details
type Student struct {
	Name string
	Age  int
}

//ShowStudent displays student details
func ShowStudent(s Student) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
}

func main() {
	s := Student{
		Name: "Muneera",
		Age:  22,
	}
	ShowStudent(s)
}
