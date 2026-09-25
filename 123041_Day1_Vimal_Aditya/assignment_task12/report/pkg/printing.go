package pkg

import "fmt"

func Title(title string) {
	fmt.Println("********************")
	fmt.Println(" ", title)
	fmt.Println("********************")
}

// Explaination: 
// pkg: It contains printing.go. This creates a public function so other folders could import and reuse.