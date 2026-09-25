// problem statement :
// ---------------------

/*
Troubleshooting 
Intentionally introduce two build/configuration errors. 
Identify the errors and document how you fixed them. 
*/

// package mains

package main


import "fmt"

func main(){

	// Println("Hello")  Error -2

	fmt.Println("Hello..")
}

/*
Error -1:
I changed package main to package mains.
The program gave build Error because the executable program must use package main
I fixed it by changing package mains back to package main
*/

/*
Error-2:
I imported the fmt package but  I didn't use it.
 so, Go gave an error "fmt" imported and not used

 I fixed it by using fmt.Println()
*/