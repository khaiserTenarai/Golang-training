package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type a message to send through the channel: ")
	messageText, _ := reader.ReadString('\n')
	message := strings.TrimSpace(messageText)

	ch := make(chan string) // unbuffered - capacity 0
	done := make(chan struct{})

	go func() {
		received := <-ch
		fmt.Println("Receiver goroutine got:", received)
		close(done)
	}()

	fmt.Println("About to send...")
	ch <- message // blocks here until the goroutine above is ready to receive
	fmt.Println("Send completed.")

	<-done // wait so the program doesn't exit before the goroutine finishes printing
}
