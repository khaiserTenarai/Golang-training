package main

import "fmt"

func main() {
	name := "ram"
	//%d is for an integer, while name is a string.
	//fmt.Printf("Employee ID: %d\n", name)
	fmt.Printf("Employee ID: %s\n", name)
}

//PS C:\Users\M.Chanikya\Downloads\Overture_Project\data\GO\Assignment1\123050_Day08_Assignment_Chanikya> go vet .\day8q6.go
//day8q6.go:8:27: fmt.Printf format %d has arg name of wrong type string
//go vet ./...
//Finds suspicious programming mistakes
