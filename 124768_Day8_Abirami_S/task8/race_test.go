package main

import (
	"sync"
	"testing"
)

func TestRace(t *testing.T) {
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter++
			}
		}()
	}
	wg.Wait()
}
