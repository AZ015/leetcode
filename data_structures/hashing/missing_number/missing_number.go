package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	sort.Ints(nums)

	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}

	return n
}

func missingNumberHashSet(nums []int) int {
	m := make(map[int]struct{})

	for _, num := range nums {
		m[num] = struct{}{}
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}

	return -1
}

func missingNumberBitManipulation(nums []int) int {
	expectedXor := 0

	for i := 0; i < len(nums); i++ {
		expectedXor ^= i
	}

	actualXor := 0
	for _, num := range nums {
		actualXor ^= num
	}

	return actualXor ^ expectedXor
}

func main() {
	fmt.Println(missingNumberBitManipulation([]int{3, 0, 1}))                   // 2
	fmt.Println(missingNumberBitManipulation([]int{0, 1}))                      // 2
	fmt.Println(missingNumberBitManipulation([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // 8
}
