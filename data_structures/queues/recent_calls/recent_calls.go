package main

import (
	"fmt"
	queues "queue"
)

func recentCalls(t int) int {
	queue := queues.New[int]()

	for !queue.IsEmpty() && queue.Peek() < t-3000 {
		queue.Pop()
	}
	queue.Push(t)

	return queue.Length()
}

func main() {
	fmt.Println(recentCalls(4000))
}
