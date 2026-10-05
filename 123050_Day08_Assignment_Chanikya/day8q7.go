package main

import "fmt"

func main() {
	x := 10
	if x == 10 {
		fmt.Println("Hello")
	}
}

/*
staticcheck .\day8q7.go
Staticcheck analyzes the code and looks for things such as:

Unnecessary code
Unreachable code
Suspicious constructs
Unused results
Inefficient code
Possible bugs
Deprecated APIs*/
