package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// produce only ever sends on ch - the compiler would reject "<-ch" here.
func produce(numbers []int, ch chan<- int) {
	for _, n := range numbers {
		ch <- n
	}
	close(ch)
}

// consume only ever receives from ch.
func consume(ch <-chan int) {
	for n := range ch {
		fmt.Println("Consumed:", n)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many numbers do you want to send? ")
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

	ch := make(chan int, count)
	go produce(numbers, ch)
	consume(ch) // runs in main, blocking until produce() closes the channel
}
