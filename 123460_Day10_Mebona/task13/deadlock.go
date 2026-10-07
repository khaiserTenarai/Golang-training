package main

func main() {
	ch := make(chan int)
	ch <- 42 // Blocks forever; no other goroutine is ready to receive.
	_ = <-ch
}