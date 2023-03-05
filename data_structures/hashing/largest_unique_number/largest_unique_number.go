package main

import (
	"fmt"
	"sort"
)

func largestUniqueNumber(nums []int) int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num] += 1
	}

	maxUniqNum := -1

	for num, count := range m {
		if count != 1 {
			continue
		}

		if maxUniqNum < num {
			maxUniqNum = num
		}
	}

	return maxUniqNum
}

func largestUniqueNumberWithSort(nums []int) int {
	sort.Ints(nums)

	for i := len(nums) - 1; i >= 0; i-- {
		if i == 0 || nums[i] != nums[i-1] {
			return nums[i]
		}

		for i > 0 && nums[i] == nums[i-1] {
			i--
		}
	}

	return -1
}

func main() {
	fmt.Println(largestUniqueNumber([]int{5, 7, 3, 9, 4, 9, 8, 3, 1}))         // 8
	fmt.Println(largestUniqueNumberWithSort([]int{5, 7, 3, 9, 4, 9, 8, 3, 1})) // 8
	fmt.Println(largestUniqueNumber([]int{9, 9, 8, 8}))                        // -1
	fmt.Println(largestUniqueNumberWithSort([]int{9, 9, 8, 8}))                // -1
}

//[8,8,9,9]
