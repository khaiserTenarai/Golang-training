package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func slowTask(durationSeconds int, done chan<- string) {
	time.Sleep(time.Duration(durationSeconds) * time.Second)
	done <- "task finished"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many seconds should the task take? ")
	taskText, _ := reader.ReadString('\n')
	taskSeconds, err := strconv.Atoi(strings.TrimSpace(taskText))
	if err != nil || taskSeconds < 0 {
		fmt.Println("Please enter a whole number, 0 or more.")
		return
	}

	fmt.Print("Timeout, in seconds: ")
	timeoutText, _ := reader.ReadString('\n')
	timeoutSeconds, err := strconv.Atoi(strings.TrimSpace(timeoutText))
	if err != nil || timeoutSeconds <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	done := make(chan string, 1)
	go slowTask(taskSeconds, done)

	select {
	case result := <-done:
		fmt.Println("Success:", result)
	case <-ctx.Done():
		fmt.Println("Timed out waiting for the task to finish.")
	}
}
