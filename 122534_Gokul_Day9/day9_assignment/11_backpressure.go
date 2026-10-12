
package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := make(chan int, 2)
	done := make(chan bool)
	start := time.Now()

	go func() {
		for t := range tasks {
			fmt.Println("   consumer: processing task", t)
			time.Sleep(700 * time.Millisecond)
		}
		done <- true
	}()

	for i := 1; i <= 6; i++ {
		fmt.Printf("[%v] producer: adding task %d\n", time.Since(start).Round(100*time.Millisecond), i)
		tasks <- i
		fmt.Printf("[%v] producer: task %d accepted\n", time.Since(start).Round(100*time.Millisecond), i)
	}
	close(tasks)

	<-done
}
