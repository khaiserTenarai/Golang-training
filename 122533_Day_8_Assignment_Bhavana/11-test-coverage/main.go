package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter marks: ")
	marksText, _ := reader.ReadString('\n')
	marks, err := strconv.Atoi(strings.TrimSpace(marksText))
	if err != nil {
		fmt.Println("Please enter a valid number.")
		return
	}

	fmt.Println("Grade:", Grade(marks))
}
