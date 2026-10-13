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

	ch := make(chan int) // unbuffered - a send needs a receiver ready right now
	ch <- num             // nobody is listening, so this blocks forever
	fmt.Println("You'll never see this line print.")
}
