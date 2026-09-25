package main

import (
	"fmt"

	"12-project-structure-simple-student-app/internal"
	"12-project-structure-simple-student-app/pkg"
)

func main() {
	student := internal.Student{
		Name: "Ganesh",
		Age: 21,
	}

	fmt.Println("Student App")
	pkg.PrintStudent(student)
}


// cmd/      → main program
// internal/ → student data/logic
// pkg/      → helper functions



/*
PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy\12-project-structure-simple-student-app> go mod init 12-project-structure-simple-student-app
go: creating new go.mod: module 12-project-structure-simple-student-app
go: to add module requirements and sums:
        go mod tidy



PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy\12-project-structure-simple-student-app> go run ./cmd
Student App
Name: Ganesh
Age: 21
		*/