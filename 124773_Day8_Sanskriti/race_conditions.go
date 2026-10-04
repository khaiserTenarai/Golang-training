package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			counter++
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}

//Run normally:

// go run race.go

// Now run with the race detector:

// go run -race race.go


// fix using mutex

//package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	counter := 0

// 	var wg sync.WaitGroup
// 	var mutex sync.Mutex

// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)

// 		go func() {
// 			defer wg.Done()

// 			mutex.Lock()
// 			counter++
// 			mutex.Unlock()
// 		}()
// 	}

// 	wg.Wait()

// 	fmt.Println("Counter:", counter)
//}