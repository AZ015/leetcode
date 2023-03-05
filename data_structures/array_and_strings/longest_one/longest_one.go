package main

import "fmt"

func longestOne(nums []int, k int) int {
	right, left := 0, 0

	for right = 0; right < len(nums); right++ {
		k -= 1 - nums[right]
		if k < 0 {
			k += 1 - nums[left]
			left += 1
		}
	}
	return right - left + 1
}

func main() {
	fmt.Println(longestOne([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}
