package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func generate(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range numbers {
			out <- n
		}
	}()
	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 2
		}
	}()
	return out
}

func addOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n + 1
		}
	}()
	return out
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many numbers do you want to send through the pipeline? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Printf("Number %d: ", i+1)
		numText, _ := reader.ReadString('\n')
		num, _ := strconv.Atoi(strings.TrimSpace(numText))
		numbers[i] = num
	}

	// each stage runs in its own goroutine, all connected by channels
	stage1 := generate(numbers)
	stage2 := double(stage1)
	stage3 := addOne(stage2)

	fmt.Println("\nPipeline output (n * 2 + 1):")
	for result := range stage3 {
		fmt.Println(result)
	}
}
