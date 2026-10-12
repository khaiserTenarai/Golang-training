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

func doWork(ctx context.Context, done chan<- struct{}) {
	defer close(done)
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Work stopped - context was cancelled.")
			return
		default:
			fmt.Println("Working... step", i)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many seconds should the work run before we cancel it? ")
	secondsText, _ := reader.ReadString('\n')
	seconds, err := strconv.Atoi(strings.TrimSpace(secondsText))
	if err != nil || seconds <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})

	go doWork(ctx, done)

	time.Sleep(time.Duration(seconds) * time.Second)
	cancel()

	<-done // wait for doWork to actually notice and stop
	fmt.Println("Main is done.")
}
