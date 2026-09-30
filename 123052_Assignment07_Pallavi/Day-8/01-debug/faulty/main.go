package main

import "fmt"

func average(nums []int) int {
	sum := 0
	for i := 0; i <= len(nums); i++ { // BUG 1: off-by-one, index out of range
		sum += nums[i]
	}
	return sum / len(nums) // BUG 2: divide by zero on an empty slice
}

func main() {
	var scores map[string]int // BUG 3: nil map
	scores["asha"] = average([]int{80, 90, 100})
	fmt.Println(scores)
}
