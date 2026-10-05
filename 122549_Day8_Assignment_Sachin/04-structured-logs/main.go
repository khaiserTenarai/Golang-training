// 4. Create structured logs.
//
// A "structured" log means the message isn't just a plain sentence, it's
// made of clear key-value fields (name, salary, etc). That makes it much
// easier for a program to search/filter logs later. Go's standard library
// has a built-in package for this since Go 1.21: log/slog.

package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func main() {
	// JSON handler prints each log line as a JSON object
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter employee name: ")
	nameText, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameText)

	fmt.Print("Enter employee salary: ")
	salaryText, _ := reader.ReadString('\n')
	salary, err := strconv.ParseFloat(strings.TrimSpace(salaryText), 64)
	if err != nil {
		logger.Error("invalid salary entered", "input", salaryText)
		return
	}

	logger.Info("employee added",
		"name", name,
		"salary", salary,
	)
}
