package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func producer(count int, ch chan<- int) {
	for i := 1; i <= count; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range ch {
		fmt.Println("Consumed item:", item)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many items should the producer make? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go producer(count, ch)
	go consumer(ch, &wg)

	wg.Wait()
	fmt.Println("Producer and consumer are both done.")
}
