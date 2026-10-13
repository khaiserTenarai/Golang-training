package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readSetting(reader *bufio.Reader, key string) (string, string) {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value, "environment variable"
	}

	fmt.Printf("%s is not set. Enter a value: ", key)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text), "typed in"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	keys := []string{"APP_NAME", "PORT", "LOG_LEVEL", "DEBUG"}
	fmt.Println()
	for _, key := range keys {
		value, source := readSetting(reader, key)
		fmt.Printf("%-10s = %-15s (%s)\n", key, value, source)
	}
}
