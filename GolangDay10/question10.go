package main

import (
	"fmt"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func squarer(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	gen := generator(2, 3, 4, 5, 6)
	sq := squarer(gen)

	for val := range sq {
		fmt.Println(val)
	}
}