// Command app is a small program used to demonstrate build automation.
package main

import "fmt"

var version = "dev"

// Greet returns a greeting for name.
func Greet(name string) string { return "Hello, " + name + "!" }

func main() {
	fmt.Println(Greet("Cloud-Native Go"), "version:", version)
}
