package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many name/score pairs do you want to add? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	var scores sync.Map
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		fmt.Printf("Name %d: ", i+1)
		nameText, _ := reader.ReadString('\n')
		name := strings.TrimSpace(nameText)

		fmt.Printf("Score for %s: ", name)
		scoreText, _ := reader.ReadString('\n')
		score, _ := strconv.Atoi(strings.TrimSpace(scoreText))

		wg.Add(1)
		go func(name string, score int) {
			defer wg.Done()
			scores.Store(name, score)
		}(name, score)
	}
	wg.Wait()

	fmt.Println("\nAll stored scores:")
	scores.Range(func(key, value interface{}) bool {
		fmt.Printf("%s -> %d\n", key, value)
		return true // returning true means "keep going to the next entry"
	})
}
