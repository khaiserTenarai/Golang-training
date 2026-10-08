package main

import (
	"fmt"
	"sync"
)

type DatabaseConnection struct{}

var (
	instance *DatabaseConnection
	once     sync.Once
)

func GetDatabaseInstance() *DatabaseConnection {
	once.Do(func() {
		fmt.Println("Initializing expensive database connection...")
		instance = &DatabaseConnection{}
	})
	return instance
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			conn := GetDatabaseInstance()
			fmt.Printf("Goroutine %d got instance %p\n", id, conn)
		}(i)
	}

	wg.Wait()
}