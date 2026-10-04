package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("Anonymous goroutine running")
	}()

	time.Sleep(time.Second)

	fmt.Println("Main finished")
}