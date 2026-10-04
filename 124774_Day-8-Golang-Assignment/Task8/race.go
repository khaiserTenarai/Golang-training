package race_detection

import "sync"

func Counter() int {
	counter := 0

	var wg sync.WaitGroup

	// Start 2 goroutines
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			// Race condition:
			// Both goroutines change counter at the same time.
			counter++
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			// Race condition:
			// Both goroutines change counter at the same time.
			counter++
		}
	}()

	wg.Wait()

	return counter
}
