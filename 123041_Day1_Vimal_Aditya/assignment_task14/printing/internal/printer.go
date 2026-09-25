package internal

import "fmt"

// lowercase 'p' makes this unexported (private)
// func printReport() {
//     fmt.Println("Generating internal report...")
// }

// correct way:
func PrintReport() {
    fmt.Println("Generating report...")
}