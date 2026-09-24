package main

import (
    "fmt"
    "github.com/google/uuid" // External dependency
)

func main() {
    id := uuid.New()
    fmt.Println("Generated ID:", id)

	/*
		PS C:\Users\mebona.joseph\Desktop\day1_go> go run configerror.go         
		configerror.go:5:5: no required module provides package github.com/google/uuid; to add it:
				go get github.com/google/uuid


		Fix: PS C:\Users\mebona.joseph\Desktop\day1_go> go mod tidy
				go: finding module for package github.com/google/uuid
				go: found github.com/google/uuid in github.com/google/uuid v1.6.0
                
				This command automatically scans the codebase for missing dependencies, 
				downloads [github.com/google/uuid](https://github.com/google/uuid), 
				and securely updates both the go.mod and go.sum files to track the correct version.
	*/
}