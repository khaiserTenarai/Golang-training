package main

import "fmt"

func main() {
	sales := [5]int{200, 300, 150, 640, 825}
	total := 0
	for i := 0; i < 5; i++ {
		total = total + sales[i]
	}
	avg := total / 5
	fmt.Println("Total Sales: ", total)
	fmt.Println("Average Sales: ", avg)
}
