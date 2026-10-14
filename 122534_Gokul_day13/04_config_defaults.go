/*
Question 4: Implement configuration defaults.

Press Enter on any question to use its default value.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ask(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func withDefault(input, fallback string) (string, bool) {
	if input == "" {
		return fallback, true
	}
	return input, false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	name, nameDefault := withDefault(ask(reader, "App name (default employee-api): "), "employee-api")
	port, portDefault := withDefault(ask(reader, "Port (default 8080): "), "8080")
	level, levelDefault := withDefault(ask(reader, "Log level (default info): "), "info")
	timeout, timeoutDefault := withDefault(ask(reader, "Timeout in seconds (default 30): "), "30")

	fmt.Println()
	show("app name", name, nameDefault)
	show("port", port, portDefault)
	show("log level", level, levelDefault)
	show("timeout", timeout, timeoutDefault)
}

func show(label, value string, usedDefault bool) {
	if usedDefault {
		fmt.Printf("%-10s = %s (default)\n", label, value)
		return
	}
	fmt.Printf("%-10s = %s\n", label, value)
}
