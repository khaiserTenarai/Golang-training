package main

import (
	"encoding/json"
	"fmt"
)

// Task 4: Structured log represented as JSON key-value fields.
func main() {
	entry := map[string]interface{}{
		"level":       "INFO",
		"message":     "Employee created",
		"employee_id": 101,
		"name":        "Ravi",
	}
	data, err := json.Marshal(entry)
	if err != nil {
		fmt.Println("Could not create log:", err)
		return
	}
	fmt.Println(string(data))
}
