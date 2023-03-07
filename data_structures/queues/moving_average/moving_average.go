package main

import (
	"fmt"
	queues "queue"
)

type MovingAverage struct {
	queue      *queues.Queue[int]
	windowSize int
	windowSum  int
}

func New(size int) *MovingAverage {
	return &MovingAverage{
		queue:      queues.New[int](),
		windowSize: size,
	}
}

func (m *MovingAverage) Next(val int) float64 {
	m.queue.Push(val)
	m.windowSum += val
	if m.queue.Length() > m.windowSize {
		m.windowSum -= m.queue.Peek()
		m.queue.Pop()
	}

	return float64(m.windowSum) / float64(m.queue.Length())
}

func main() {
	average := New(3)
	fmt.Println(average.Next(1))  //1.0
	fmt.Println(average.Next(10)) //5.5
	fmt.Println(average.Next(3))  //4.666667
	fmt.Println(average.Next(5))  //6.0

}
