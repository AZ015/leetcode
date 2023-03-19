package main

import (
	"container/heap"
	"fmt"
	"heaps"
	"math"
)

func lastStoneWeight(stones []int) int {
	h := &heaps.IntHeap{}
	heap.Init(h)
	for _, stone := range stones {
		heap.Push(h, stone)
	}

	for h.Len() > 1 {
		first := int(math.Abs(float64(h.Pop().(int))))
		second := int(math.Abs(float64(h.Pop().(int))))
		if first != second {
			h.Push(first - second)
		}
	}

	if h.Len() == 0 {
		return 0
	}

	return h.Pop().(int)
}

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}
