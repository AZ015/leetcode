package main

import (
	"fmt"
	"math"
	queues "queue"
)

func longestSubarray(nums []int, limit int) int {
	increasing := queues.New[int]()
	decreasing := queues.New[int]()
	left, ans := 0, 0

	for right := 0; right < len(nums); right++ {
		for !increasing.IsEmpty() && increasing.Front() > nums[right] {
			increasing.Pop()
		}
		for !decreasing.IsEmpty() && decreasing.Front() < nums[right] {
			decreasing.Pop()
		}
		increasing.Push(nums[right])
		decreasing.Push(nums[right])

		for decreasing.Back()-increasing.Back() > limit {
			if nums[left] == decreasing.Back() {
				decreasing.PopLeft()
			}
			if nums[left] == increasing.Back() {
				increasing.PopLeft()
			}
			left += 1
		}
		ans = int(math.Max(float64(ans), float64(right-left+1)))
	}

	return ans
}

func main() {
	fmt.Println(longestSubarray([]int{10, 1, 2, 4, 7, 2}, 5))
}
