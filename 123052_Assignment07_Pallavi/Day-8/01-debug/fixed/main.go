package main

import (
	"errors"
	"fmt"
)

func average(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("cannot average an empty slice")
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum / len(nums), nil
}

func main() {
	scores := make(map[string]int)
	avg, err := average([]int{80, 90, 100})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	scores["asha"] = avg
	fmt.Println(scores)
}
