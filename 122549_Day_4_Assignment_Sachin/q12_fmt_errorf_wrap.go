// 12. Use fmt.Errorf with %w.

package main

import (
	"errors"
	"fmt"
)

var ErrDatabase = errors.New("connection failed")

func fetchRecord() error {
	return ErrDatabase
}

func processData() error {
	err := fetchRecord()
	if err != nil {
		return fmt.Errorf("processData failed: %w", err)
	}
	return nil
}

func main() {
	err := processData()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
