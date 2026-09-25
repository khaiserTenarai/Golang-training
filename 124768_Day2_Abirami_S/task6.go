package main

import "fmt"

func main() {
	var attendance int
	var marks int
	fmt.Println("Enter your attendance: ")
	fmt.Scan(&attendance)
	fmt.Println("Enter your marks: ")
	fmt.Scan(&marks)
	if marks >= 90 && attendance >= 95 {
		fmt.Println("Student is eligible for 90% Scholarship")
	} else {
		fmt.Println("Student is not eligible for 90% Scholarship")
	}
	if marks >= 95 || attendance == 100 {
		fmt.Println("Student is eligible for 100% Scholarship")
	}
}
