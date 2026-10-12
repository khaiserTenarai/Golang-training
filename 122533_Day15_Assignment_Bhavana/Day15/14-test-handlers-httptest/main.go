package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter port number (press Enter for 8080): ")
	text, _ := reader.ReadString('\n')
	port := strings.TrimSpace(text)
	if port == "" {
		port = "8080"
	}

	router := NewRouter()
	if err := router.Run(":" + port); err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
