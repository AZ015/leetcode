package main

import "fmt"

func countElements(arr []int) int {
	m := make(map[int]struct{})

	for _, num := range arr {
		m[num] = struct{}{}
	}

	cnt := 0

	for _, num := range arr {
		if _, ok := m[num+1]; ok {
			cnt += 1
		}
	}

	return cnt
}

func main() {
	fmt.Println(countElements([]int{1, 2, 3}))
	fmt.Println(countElements([]int{1, 1, 3, 3, 5, 5, 7, 7}))
}
