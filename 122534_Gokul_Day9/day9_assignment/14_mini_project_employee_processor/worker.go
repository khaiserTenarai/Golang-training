package main

import (
	"sync"
	"time"
)

func worker(id int, jobs <-chan Employee, results chan<- Payslip, wg *sync.WaitGroup) {
	defer wg.Done()

	for emp := range jobs {
		time.Sleep(300 * time.Millisecond)
		results <- calculatePayslip(emp, id)
	}
}
