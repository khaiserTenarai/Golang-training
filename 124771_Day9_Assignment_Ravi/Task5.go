package main

import "fmt"

func main() {

	ch := make(chan string, 3)

	ch <- "Order-101"
	ch <- "Order-102"
	ch <- "Order-103"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
