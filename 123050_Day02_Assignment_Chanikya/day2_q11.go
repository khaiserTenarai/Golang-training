package main

import "fmt"

func main() {

	libraryCatalog := map[int]string{
		9001: "The Go Programming Language",
		9002: "Clean Code",
		9003: "Introduction to Algorithms",
	}

	var isbnCode int

	fmt.Print("Enter Book ISBN Code: ")
	fmt.Scan(&isbnCode)

	title, isAvailable := libraryCatalog[isbnCode]

	if isAvailable {
		fmt.Println("Book Title:", title)
	} else {
		fmt.Println("Book not found")
	}
}
