// Command versioning prints its version, which is injected at build time:
//
//	go build -ldflags "-X main.version=$(git describe --tags)" .
package main

import "fmt"

// version is overwritten by -ldflags at build time.
var version = "dev"

func main() {
	fmt.Println("versioning demo, version:", version)
}
