// 8. Go Documentation
//
// Add GoDoc comments to exported functions and structures.
// Generate package documentation.

package main

import (
	"fmt"
	"godocdemo/mathutils"
)

func main() {
	fmt.Println("Sum:", mathutils.Add(4, 5))
	fmt.Println("Difference:", mathutils.Subtract(9, 3))
}
