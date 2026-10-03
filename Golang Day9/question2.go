package main

import (
	"fmt"
	"sync"
	"time"
)

func calculateBonus(empID int,wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Starting calculation empId",empID)
	time.Sleep(2 *time.Second)
	fmt.Println("Finished the clacaulation of empId",empID)

}
func main() {
	var wg sync.WaitGroup
	employees := []int{101,102,103,104}
	fmt.Println("Starting Payroll")
	for _,empID := range employees {
		wg.Add(1)
		go calculateBonus(empID,&wg)
	}
	wg.Wait()
	fmt.Println("All the employee calculations are complete")
}