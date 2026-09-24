package main

import "fmt"

func main() {
	name := "Pallavi"

	fmt.Println("Hello,", name)
}


// 1. go run(Run the program)-- Compiles and runs the Go program directly without creating a permanent executable.
// 2. go build(Build the program)-- Compiles the Go program and creates an executable file.
// 3. go fmt(Format the code)-- Automatically formats Go code according to Go's standard formatting style.
// 4. go vet(Check the code)-- Finds common mistakes and suspicious code that may cause problems.
// 5. go test(Run tests)-- Runs test functions written in Go test files (*_test.go) and checks whether the tests pass.