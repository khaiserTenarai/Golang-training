package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	AppName string
	Port    int
	Debug   bool
}

func ask(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := ask(reader, "Enter the app name: ")
	port, err := strconv.Atoi(ask(reader, "Enter the port: "))
	if err != nil {
		fmt.Println("port must be a number")
		return
	}
	debug := strings.ToLower(ask(reader, "Turn on debug mode? (y/n): ")) == "y"

	cfg := Config{AppName: name, Port: port, Debug: debug}

	fmt.Printf("%+v\n", cfg)
	fmt.Println("Starting", cfg.AppName, "on port", cfg.Port)
}
