package main

import (
    "fmt"
    "os"
    "runtime"
)

func main() {
    fmt.Println("Checking the Go installation on this machine...")

    fmt.Println("Go Version ->", runtime.Version())

    fmt.Println("GOPATH     ->", os.Getenv("GOPATH"))
}