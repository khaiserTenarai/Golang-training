package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello from goroutine")
}

func main() {
	go hello()
	time.Sleep(time.Second)
	fmt.Println("Main complete")
}
