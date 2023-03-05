package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	var sum float64 = 0
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}

	var res = sum
	for i := k; i < len(nums); i++ {
		sum += float64(nums[i] - nums[i-k])
		res = math.Max(res, sum)
	}

	return res / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))
}
