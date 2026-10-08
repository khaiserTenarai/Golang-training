package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex1 sync.Mutex
	var mutex2 sync.Mutex

	var wg sync.WaitGroup

	wg.Add(2)

	// go func() {
	// 	defer wg.Done()

	// 	mutex1.Lock()
	// 	fmt.Println("Goroutine 1 locked mutex1")

	// 	mutex2.Lock()
	// 	fmt.Println("Goroutine 1 locked mutex2")

	// 	mutex2.Unlock()
	// 	mutex1.Unlock()
	// }()

	// go func() {
	// 	defer wg.Done()

	// 	mutex2.Lock()
	// 	fmt.Println("Goroutine 2 locked mutex2")

	// 	mutex1.Lock()
	// 	fmt.Println("Goroutine 2 locked mutex1")

	// 	mutex1.Unlock()
	// 	mutex2.Unlock()
	// }()

	// Fix - acquire lock in the same order
	go func() {
		defer wg.Done()

		mutex1.Lock()
		defer mutex1.Unlock()

		mutex2.Lock()
		defer mutex2.Unlock()

		fmt.Println("Goroutine 1 completed")
	}()

	go func() {
		defer wg.Done()

		mutex1.Lock()
		defer mutex1.Unlock()

		mutex2.Lock()
		defer mutex2.Unlock()

		fmt.Println("Goroutine 2 completed")
	}()
	wg.Wait()

	fmt.Println("Completed")

}
