package main

import (
	"bufio"
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

	name := ask(reader, "Enter a name for the worker: ")
	if name == "" {
		name = "worker"
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	fmt.Printf("%s started (pid %d). Press Ctrl+C to stop\n", name, os.Getpid())

	jobs := 0
	for {
		select {
		case <-ticker.C:
			jobs++
			fmt.Printf("%s: finished job %d\n", name, jobs)
		case sig := <-signals:
			fmt.Printf("\nreceived %v, cleaning up\n", sig)
			fmt.Printf("%s stopped after %d jobs\n", name, jobs)
			return
		}
	}
}
