package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		Two goroutines are changing the same variable.

		This creates a race condition.
	*/

	var count int
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// ❌ Shared data without protection.
			count++
		}()
	}

	wg.Wait()

	fmt.Println("Count:", count)
}
/*
PS C:\Training\Assignments\Day_10\14-Find and Fix a Race Condition Using -race> go run -race main.go
==================
WARNING: DATA RACE
Read at 0x00c00008e078 by goroutine 9:
  main.main.func1()
      C:/Training/Assignments/Day_10/14-Find and Fix a Race Condition Using -race/main.go:25 +0x7b

Previous write at 0x00c00008e078 by goroutine 8:
  main.main.func1()
      C:/Training/Assignments/Day_10/14-Find and Fix a Race Condition Using -race/main.go:25 +0x8d

Goroutine 9 (running) created at:
  main.main()
      C:/Training/Assignments/Day_10/14-Find and Fix a Race Condition Using -race/main.go:21 +0x78

Goroutine 8 (finished) created at:
  main.main()
      C:/Training/Assignments/Day_10/14-Find and Fix a Race Condition Using -race/main.go:21 +0x78
==================
Count: 2
Found 1 data race(s)
exit status 66
*/