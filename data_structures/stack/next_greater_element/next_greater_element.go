package main

import (
	"fmt"
	. "stack"
)

func nextGreaterElement(nums1, nums2 []int) []int {
	res := make([]int, len(nums1))
	stack := NewStackGeneric[int]()
	m := make(map[int]int)

	for _, num := range nums2 {
		for !stack.IsEmpty() && num > stack.Peek() {
			m[stack.Pop()] = num
		}
		stack.Push(num)
	}

	for !stack.IsEmpty() {
		m[stack.Pop()] = -1
	}

	for i, num := range nums1 {
		res[i] = m[num]
	}

	return res
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})) // [-1, 3, -1]
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))    // [3, -1]
}
