package utils

import "fmt"

func PrintHeader(title string) {
	fmt.Println("---", title, "---")
}

func PrintFooter(message string) {
	fmt.Println("---", message, "---")
}

func Bonus(salary int) int {
	return salary + 5000
}