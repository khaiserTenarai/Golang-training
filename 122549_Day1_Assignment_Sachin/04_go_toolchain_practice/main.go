// Day 1 - Task 4: Go Toolchain Practice
//
//
//   go run main.go   compiles and runs the program in one shot without
//                     leaving a binary behind. Handy for quick checks.
//
//   go build         compiles the package into a standalone binary
//                     (e.g. ./toolchain) that can run without Go
//                     installed on the target machine.
//
//   go fmt ./...     reformats the source to Go's standard style
//                     (gofmt), so spacing/indentation stays consistent.
//
//   go vet ./...     statically checks the code for constructs that
//                     compile fine but are probably bugs (e.g. a wrong
//                     Printf verb).
//
//   go test ./...    runs every *_test.go file and reports pass/fail.
package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	sum := Add(4, 5)
	fmt.Println("4 + 5 =", sum)
}
