package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if _, ok := m[complement]; ok {
			return []int{m[num], i}
		}
		m[num] = i
	}

	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{5, 2, 7, 10, 3, 9}, 8))
}
