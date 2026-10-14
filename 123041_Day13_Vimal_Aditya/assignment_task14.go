package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Features struct {
	NewBonusRule bool
}

func ask(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func loadFeatures(reader *bufio.Reader) Features {
	if value, found := os.LookupEnv("FEATURE_NEW_BONUS"); found {
		return Features{NewBonusRule: value == "true"}
	}
	answer := ask(reader, "Turn on the new bonus rule? (y/n): ")
	return Features{NewBonusRule: strings.ToLower(answer) == "y"}
}

func calculateBonus(salary float64, features Features) float64 {
	if features.NewBonusRule {
		return salary * 0.10
	}
	return salary * 0.05
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	features := loadFeatures(reader)

	name := ask(reader, "Enter the employee name: ")
	salary, err := strconv.ParseFloat(ask(reader, "Enter the salary: "), 64)
	if err != nil {
		fmt.Println("salary must be a number")
		return
	}

	bonus := calculateBonus(salary, features)
	fmt.Printf("%s gets a bonus of %.2f (new rule on: %v)\n", name, bonus, features.NewBonusRule)
}
