
package main

import "troubleshooting/internal"

// Error 1:
// cannot refer to unexported name internal.printReport
//
// func main() {
//     internal.printReport()
// }

// Correct way:
func main() {
	internal.PrintReport()
}


