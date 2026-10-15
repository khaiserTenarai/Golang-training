package main

import "fmt"

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func addTen(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n + 10
		}
	}()
	return out
}

func main() {
	nums := generator(2, 3, 4)
	squared := square(nums)
	result := addTen(squared)

	for res := range result {
		fmt.Println("Pipeline result:", res)
	}
}