package main

import "printing/internal"

// below code will throw error -
// cannot refer to unexported name internal.printReport
// func main() {
//     internal.printReport() // Attempting to access private function
// }

// correct way:
func main() {
    internal.PrintReport() // Attempting to access private function
}