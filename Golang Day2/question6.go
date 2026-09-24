package main

import "fmt"

func main() {

	var age,height int
	fmt.Println("Enter your name:")
    fmt.Scan(&age)
    fmt.Println("Enter your height in cm:")
    fmt.Scan(&height)
    
	if age>=18 && height>=120 {
		fmt.Println("You are selected")
	}

	if age>=17 || height<120 {
		fmt.Println("You are not selected")
	}

}