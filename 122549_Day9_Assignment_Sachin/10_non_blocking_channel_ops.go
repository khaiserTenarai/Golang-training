package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to put a message in the channel first? (yes/no): ")
	answerText, _ := reader.ReadString('\n')
	answer := strings.ToLower(strings.TrimSpace(answerText))

	ch := make(chan string, 1)

	if answer == "yes" {
		fmt.Print("Enter the message: ")
		msgText, _ := reader.ReadString('\n')
		ch <- strings.TrimSpace(msgText)
	}

	// This select never blocks - if ch has nothing to give us right
	// now, it falls through to default instead of waiting.
	select {
	case msg := <-ch:
		fmt.Println("Got a message:", msg)
	default:
		fmt.Println("No message was waiting, moved on immediately.")
	}
}
