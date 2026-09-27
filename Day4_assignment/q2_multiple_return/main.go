
// Day 4, Q2. Create a function returning multiple values.

package main

import "fmt"

func splitName(fullName string) (first string, last string) {
	space := -1
	for i, ch := range fullName {
		if ch == ' ' {
			space = i
			break
		}
	}
	if space == -1 {
		return fullName, ""
	}
	return fullName[:space], fullName[space+1:]
}

func minMax(nums []int) (min int, max int) {
	min, max = nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}

func main() {
	first, last := splitName("Ranjitha ray")
	fmt.Println("First:", first, "Last:", last)

	smallest, largest := minMax([]int{45, 12, 89, 3, 67})
	fmt.Println("Min:", smallest, "Max:", largest)
}
