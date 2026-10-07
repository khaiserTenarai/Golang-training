package main

import (
	"fmt"
	"sync"
)

type SafeDictionary struct {
	mu   sync.RWMutex
	data map[string]string
}

func (d *SafeDictionary) Write(key, value string, wg *sync.WaitGroup) {
	defer wg.Done()
	
	d.mu.Lock()
	d.data[key] = value
	d.mu.Unlock()
}

func (d *SafeDictionary) Read(key string, wg *sync.WaitGroup) string {
	defer wg.Done()
	
	d.mu.RLock()
	val := d.data[key]
	d.mu.RUnlock()
	
	return val
}

func main() {
	dict := SafeDictionary{
		data: make(map[string]string),
	}
	var wg sync.WaitGroup

	wg.Add(2)
	go dict.Write("db_host", "localhost", &wg)
	go dict.Write("db_port", "5432", &wg)

	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			dict.Read("db_host", &wg)
		}()
	}

	wg.Wait()
	
	fmt.Println("Config loaded:")
	fmt.Println("Host:", dict.data["db_host"])
	fmt.Println("Port:", dict.data["db_port"])
}