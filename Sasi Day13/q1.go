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

func envOrAsk(reader *bufio.Reader, key string) string {
	if value := os.Getenv(key); value != "" {
		fmt.Printf("%s found in the environment\n", key)
		return value
	}
	return ask(reader, fmt.Sprintf("%s is not set, enter a value: ", key))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := envOrAsk(reader, "APP_NAME")
	port := envOrAsk(reader, "APP_PORT")
	fmt.Println("APP_NAME =", name)
	fmt.Println("APP_PORT =", port)

	key := ask(reader, "Enter any environment variable to look up (for example HOME): ")
	if value, found := os.LookupEnv(key); found {
		fmt.Printf("%s = %s\n", key, value)
	} else {
		fmt.Printf("%s is not set\n", key)
	}
}
