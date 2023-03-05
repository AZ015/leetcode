package main

import "fmt"

func runningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	prefix := nums[0]
	for i := 1; i < len(nums); i++ {
		prefix += nums[i]
		nums[i] = prefix
	}

	return nums
}

func runningSumFinal(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return nums
}

func main() {
	fmt.Println(runningSumFinal([]int{1, 2, 3, 4}))     // [1,3,6,10]
	fmt.Println(runningSumFinal([]int{1, 1, 1, 1, 1}))  // [1,2,3,4,5]
	fmt.Println(runningSumFinal([]int{3, 1, 2, 10, 1})) // [3,4,6,16,17]
}
