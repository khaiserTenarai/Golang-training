package main

import "fmt"

func idGenerator() func() int {
	id := 100
	return func() int {
		id++
		return id
	}
}
func main() {
	id := idGenerator()
	fmt.Println("Employee1 ID: ", id())
	fmt.Println("Employee2 ID: ", id())
	fmt.Println("Employee3 ID: ", id())
	fmt.Println("Employee4 ID: ", id())
}
