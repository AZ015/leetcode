package main

import (
	"fmt"
	queues "queue"
)

func slidingWindowMaximum(arr []int, k int) []int {
	ans := make([]int, 0)
	queue := queues.New[int]()

	for i, el := range arr {
		for !queue.IsEmpty() && el > arr[queue.Back()] {
			queue.Pop()
		}
		queue.Push(i)

		if queue.Front()+k == i {
			queue.PopLeft()
		}
		if i > k-1 {
			ans = append(ans, arr[queue.Front()])
		}
	}

	return ans
}

func main() {
	fmt.Println(slidingWindowMaximum([]int{1, 3, -1, -3, -2, 3, 6, 7}, 3))
}
