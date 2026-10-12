package main

import "fmt"

func main() {
	salaries := make(chan int)

	go func() {
		for _, s := range []int{30000, 45000, 38000, 52000} {
			salaries <- s
		}
		close(salaries)
	}()

	total := 0
	for s := range salaries {
		fmt.Println("Got salary:", s)
		total += s
	}
	fmt.Println("Total salary:", total)
}
