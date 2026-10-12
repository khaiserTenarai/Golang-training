package main

import "fmt"

func consumer(results <-chan Payslip, done chan<- Summary) {
	var summary Summary

	for slip := range results {
		fmt.Printf("consumer: %-8s (%-7s) net salary %.2f  [worker %d]\n",
			slip.Employee.Name, slip.Employee.Department, slip.Net, slip.WorkerID)
		summary.Count++
		summary.TotalNet += slip.Net
	}

	done <- summary
}
