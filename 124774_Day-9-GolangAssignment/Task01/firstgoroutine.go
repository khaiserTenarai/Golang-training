package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine")
}

func main() {
	go sayHello()

	time.Sleep(time.Second)
	fmt.Println("Main function finished")
}
