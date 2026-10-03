package main

import "fmt"

func main() {
	name := "Alice"

	fmt.Printf("Employee ID: %d\n", name)
}

/*
PS C:\Training\Go Lang\Day_9\124772_Day_9(Go)_Assignment_Reddem_Ganesh_Reddy\06-govet> go mod init govet
go: creating new go.mod: module govet
go: to add module requirements and sums:
        go mod tidy
PS C:\Training\Go Lang\Day_9\124772_Day_9(Go)_Assignment_Reddem_Ganesh_Reddy\06-govet> go vet

main.go:8:27: fmt.Printf format %d has arg name of wrong type string

*/
