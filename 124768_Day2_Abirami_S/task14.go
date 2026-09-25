package main

import "fmt"

var name = "global"

func main() {
	fmt.Println("Global Scope: ", name)
	name := "local"
	fmt.Println("Local Scope: ", name)
	{
		name := "Shadowing"
		fmt.Println("Block level access: ", name)
	}
	fmt.Println("Local Scope after block: ", name)
}
