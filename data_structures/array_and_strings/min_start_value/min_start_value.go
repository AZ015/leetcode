package main

import (
	"fmt"
)

func minStartValue(nums []int) int {
	minValue := 0
	total := 0

	for _, num := range nums {
		total += num
		minValue = min(minValue, total)
	}

	return -minValue + 1
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}

	return num2
}

func main() {
	fmt.Println(minStartValue([]int{1, 2}))            // 1
	fmt.Println(minStartValue([]int{1, -2, -3}))       // 5
	fmt.Println(minStartValue([]int{-3, 2, -3, 4, 2})) // 5
}
