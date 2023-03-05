package main

import (
	"fmt"
	"sort"
)

func maxNumbersOfBalloons(text string) int {
	bCount, aCount, lCount, oCount, nCount := 0, 0, 0, 0, 0

	for i := 0; i < len(text); i++ {
		switch string(text[i]) {
		case "b":
			bCount += 1
		case "a":
			aCount += 1
		case "l":
			lCount += 1
		case "o":
			oCount += 1
		case "n":
			nCount += 1
		}
	}

	lCount /= 2
	oCount /= 2

	return min([]int{aCount, bCount, aCount, oCount, lCount, nCount})
}

func min(nums []int) int {
	sort.Ints(nums)

	return nums[0]
}

func main() {
	fmt.Println(maxNumbersOfBalloons("nlaebolko"))        // 1
	fmt.Println(maxNumbersOfBalloons("loonbalxballpoon")) // 2
	fmt.Println(maxNumbersOfBalloons("leetcode"))         // 0
}
