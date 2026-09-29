// 2. Create a function returning multiple values.

package main

import "fmt"

func minMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func main() {
	scores := []int{45, 89, 12, 67, 90, 33}

	min, max := minMax(scores)
	fmt.Println("Scores:", scores)
	fmt.Println("Minimum:", min)
	fmt.Println("Maximum:", max)
}
