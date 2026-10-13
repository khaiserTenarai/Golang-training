package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number to send into the channel: ")
	numText, _ := reader.ReadString('\n')
	num, err := strconv.Atoi(strings.TrimSpace(numText))
	if err != nil {
		fmt.Println("Please enter a whole number.")
		return
	}

	ch := make(chan int)
	done := make(chan struct{})

	go func() {
		received := <-ch
		fmt.Println("Received:", received)
		close(done)
	}()

	ch <- num // this works now - the goroutine above is ready to receive
	<-done    // wait for it to finish printing before we exit
	fmt.Println("No deadlock this time.")
}
