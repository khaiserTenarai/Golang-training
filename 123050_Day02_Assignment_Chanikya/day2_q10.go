package main

import "fmt"

type Student struct {
	RollNo int
	Name   string
}

func main() {
	students := []Student{}

	students = append(students, Student{501, "Amit"})
	students = append(students, Student{502, "Kiran"})

	fmt.Println("Student Records:")
	for _, student := range students {
		fmt.Println(student.RollNo, student.Name)
	}

	students[0].Name = "Amit Sharma"

	students = students[:1]

	fmt.Println("\nAfter Upgrading and Removing:")
	for _, student := range students {
		fmt.Println(student.RollNo, student.Name)
	}
}
