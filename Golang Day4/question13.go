package main

import (
	"errors"
	"fmt"
)

var ErrPermissionDenied = errors.New("permission denied")

func deleteFile(filename string) error {
	return fmt.Errorf("cannot delete %s: %w", filename, ErrPermissionDenied)
}

func main() {
	err := deleteFile("system_config.sys")

	if err != nil {
		fmt.Println("Error:", err)

		if errors.Is(err, ErrPermissionDenied) {
			fmt.Println("Run the program as administrator")
		}
	}
}