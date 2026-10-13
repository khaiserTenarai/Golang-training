package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func settingOrDefault(reader *bufio.Reader, key, fallback string) (string, string) {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value, "environment variable"
	}

	fmt.Printf("%s is not set (Enter to use the default %q): ", key, fallback)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		return fallback, "default"
	}
	return text, "typed in"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	defaults := []struct{ key, value string }{
		{"APP_NAME", "employee-api"},
		{"PORT", "8080"},
		{"LOG_LEVEL", "info"},
		{"REQUEST_TIMEOUT_SECONDS", "30"},
	}

	fmt.Println()
	for _, d := range defaults {
		value, source := settingOrDefault(reader, d.key, d.value)
		fmt.Printf("  -> %s = %s  [%s]\n\n", d.key, value, source)
	}
}
