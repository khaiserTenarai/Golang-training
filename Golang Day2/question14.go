package main

import "fmt"

// 1. GLOBAL SCOPE
// This variable lives outside all functions. 
// Everyone in this file can see and use it.
var hero = "Batman"

func main() {
	fmt.Println("1. Global scope   :", hero) 

	// 2. LOCAL SCOPE (SHADOWING)
	// Because it is inside main(), it "shadows" (hides) the global Batman.
	hero := "Spider-Man"
	fmt.Println("2. Local scope    :", hero)

	// 3. BLOCK SCOPE (DEEP SHADOWING)
	// This 'if' statement has its own set of curly braces { }
	if true {
		// We create ANOTHER new variable named 'hero' just for this block.
		hero := "Iron Man"
		fmt.Println("3. Block scope    :", hero) 
	}

	// 4. BACK TO LOCAL
	// The 'if' block ended, so "Iron Man" was destroyed. 
	// The shadow is gone and we see the local main() variable again.
	fmt.Println("4. Back to Local  :", hero) 
}