package main
import (
	"fmt"
	"sync"
	"time"
)

/*
	GOROUTINE LIFECYCLE

	1. Create   → using "go"
	2. Run      → goroutine executes
	3. Wait     → goroutine waits for some time
	4. Resume   → goroutine continues
	5. Finish   → function ends
*/

func task(wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Goroutine started")

	// Goroutine waits for 1 second
	time.Sleep(1 * time.Second)

	fmt.Println("Goroutine resumed")

	fmt.Println("Goroutine finished")
}

func main() {

	var wg sync.WaitGroup

	// We are waiting for one goroutine
	wg.Add(1)

	fmt.Println("Main started")

	// Create a goroutine
	go task(&wg)

	// Wait until goroutine finishes
	wg.Wait()

	fmt.Println("Main finished")
}

/*
Output :
----------
PS C:\Training\HP - GO Language Training\Day_9\124772_Day_9(Go)_Assignment_Reddem_Ganesh_Reddy\03-Go_Routine_Life_Cycle> go run .\main.go
Main started
Goroutine started
Goroutine resumed
Goroutine finished
Main finished
*/