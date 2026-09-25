// 4. Go Toolchain Practice
//
// Execute:
// - go run
// - go build
// - go fmt
// - go vet
// - go test

package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("2 + 3 =", Add(2, 3))
}
