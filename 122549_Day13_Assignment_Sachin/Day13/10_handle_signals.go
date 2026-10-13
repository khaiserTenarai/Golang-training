package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many seconds between work ticks? ")
	text, _ := reader.ReadString('\n')
	seconds, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil || seconds <= 0 {
		fmt.Println("Not a valid number, using 2.")
		seconds = 2
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	ticker := time.NewTicker(time.Duration(seconds) * time.Second)
	defer ticker.Stop()

	fmt.Println("Working... press Ctrl+C to stop.")
	for {
		select {
		case <-ticker.C:
			fmt.Println("  tick - doing some work")
		case sig := <-signals:
			fmt.Println("\nReceived signal:", sig)
			fmt.Println("Saving state and cleaning up...")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Goodbye.")
			return
		}
	}
}
