package main

import (
	"fmt"
	"sort"
)

func sortedSquaresFirst(nums []int) []int {
	for i, num := range nums {
		nums[i] = num * num
	}

	sort.Ints(nums)

	return nums
}

func sortedSquaresFinal(nums []int) []int {
	n := len(nums)
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1

	for i := n - 1; i >= 0; i-- {
		lv, rv := nums[left]*nums[left], nums[right]*nums[right]
		if lv < rv {
			res[i] = rv
			right--
		} else {
			res[i] = lv
			left++
		}
	}

	return res
}

func main() {
	fmt.Println(sortedSquaresFinal([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquaresFinal([]int{-7, -3, -2, 3, 11}))
}
