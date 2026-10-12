
package main

import (
	"fmt"
	"time"
)

func main() {
	payroll := make(chan string)
	attendance := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		payroll <- "Payroll report is ready"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		attendance <- "Attendance report is ready"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-payroll:
			fmt.Println(msg)
		case msg := <-attendance:
			fmt.Println(msg)
		}
	}

	nothing := make(chan string)
	select {
	case msg := <-nothing:
		fmt.Println(msg)
	case <-time.After(1500 * time.Millisecond):
		fmt.Println("Timed out, no message arrived")
	}
}
