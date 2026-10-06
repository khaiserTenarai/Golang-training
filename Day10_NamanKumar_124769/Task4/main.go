package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			once.Do(func() {
				fmt.Println("init runs once , called by", id)
			})
		}(i)
	}
	wg.Wait()
}
