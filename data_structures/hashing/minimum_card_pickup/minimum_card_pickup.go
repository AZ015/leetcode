package main

import (
	"fmt"
	"math"
)

func minimumCardPickup(cards []int) int {
	m := make(map[int][]int)
	ans := math.MaxInt

	for i, card := range cards {
		if _, ok := m[card]; ok {
			ans = int(math.Min(float64(ans), float64(i-m[card][i]+1)))
		}
		m[card] = append(m[card], i)
	}

	if ans == math.MaxInt {
		return -1
	}

	return ans
}

func main() {
	fmt.Println(minimumCardPickup([]int{1, 2, 6, 2, 1}))
}
