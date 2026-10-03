package main

import(
	"fmt"
	"sync"
) 

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		defer wg.Done()
		fmt.Println("Hello")
		fmt.Println("Employee processing started")
	}()
	wg.Wait()
	fmt.Println("Main function completed")
}
