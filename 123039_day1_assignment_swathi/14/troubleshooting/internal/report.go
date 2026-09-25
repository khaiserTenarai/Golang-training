
package internal

import "fmt"

// Error 2:
// printReport starts with lowercase 'p',
// so it is unexported.
//
// func printReport() {
//     fmt.Println("Generating internal report...")
// }

// Correct way:
func PrintReport() {
	fmt.Println("Generating report...")
}


