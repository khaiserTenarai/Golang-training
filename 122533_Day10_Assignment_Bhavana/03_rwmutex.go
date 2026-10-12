package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	price   float64
	rwMutex sync.RWMutex
)

func readPrice(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	rwMutex.RLock()
	fmt.Printf("Reader %d sees price: %.2f\n", id, price)
	rwMutex.RUnlock()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Starting price: ")
	startText, _ := reader.ReadString('\n')
	startPrice, err := strconv.ParseFloat(strings.TrimSpace(startText), 64)
	if err != nil {
		fmt.Println("Please enter a valid number.")
		return
	}
	price = startPrice

	fmt.Print("How many readers should check the price? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	var wg sync.WaitGroup
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go readPrice(i, &wg)
	}
	wg.Wait()

	fmt.Print("New price: ")
	newText, _ := reader.ReadString('\n')
	newPrice, err := strconv.ParseFloat(strings.TrimSpace(newText), 64)
	if err != nil {
		fmt.Println("Please enter a valid number.")
		return
	}

	rwMutex.Lock()
	price = newPrice
	rwMutex.Unlock()
	fmt.Println("Price updated to:", price)
}
