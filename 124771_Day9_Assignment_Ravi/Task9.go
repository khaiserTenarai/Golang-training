package main

import (
	"fmt"

	"time"
)

func main() {

	ch := make(chan string)

	go func() {

		time.Sleep(500 * time.Millisecond)

		ch <- "Data received"

	}()

	select {

	case msg := <-ch:

		fmt.Println(msg)

	case <-time.After(time.Second):

		fmt.Println("Timed out")

	}

}
