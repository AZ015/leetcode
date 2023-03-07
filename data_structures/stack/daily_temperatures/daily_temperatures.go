package main

import (
	"fmt"
)

/*
	Given an array of integers temperatures that represent the daily temperatures, return an array answer such that
	answer[i] is the number of days you have to wait after the i day to get a warmer temperature. If there is no
	future day that is warmer, have answer[i] = 0 instead
*/

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	answer := make([]int, len(temperatures))

	for i, _ := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[j] = i - j
		}
		stack = append(stack, i)
	}

	return answer
}

func main() {
	fmt.Println(dailyTemperatures([]int{40, 35, 32, 37, 50}))
}
