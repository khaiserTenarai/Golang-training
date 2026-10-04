package main

import (
	"errors"
	"fmt"
)

type NetworkError struct {
	Code    int
	Message string
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("network error %d: %s", e.Code, e.Message)
}

func fetchData() error {
	return fmt.Errorf("fetch failed: %w", &NetworkError{Code: 503, Message: "service unavailable"})
}

func main() {
	err := fetchData()

	if err != nil {
		fmt.Println("Error:", err)

		var netErr *NetworkError
		if errors.As(err, &netErr) {
			fmt.Println("Extracted HTTP Code:", netErr.Code)
			fmt.Println("Extracted Message:", netErr.Message)
		}
	}
}