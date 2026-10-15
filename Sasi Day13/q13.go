package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func ask(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	list := ask(reader, "Enter the resources to start, separated by commas (for example logger,database,server): ")

	var started []string
	for _, name := range strings.Split(list, ",") {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		fmt.Println("starting", name)
		started = append(started, name)
	}
	if len(started) == 0 {
		fmt.Println("nothing to start")
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	enter := make(chan struct{})
	go func() {
		reader.ReadString('\n')
		close(enter)
	}()

	fmt.Println("Everything is running. Press Enter or Ctrl+C to shut down")
	select {
	case <-ctx.Done():
	case <-enter:
	}

	fmt.Println("\nshutdown started")
	for i := len(started) - 1; i >= 0; i-- {
		fmt.Println("stopping", started[i])
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("shutdown complete")
}
