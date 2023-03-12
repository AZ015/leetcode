package trees_and_graphs

type Queue[T any] struct {
	Queue []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Queue: make([]T, 0),
	}
}

func (q *Queue[T]) PeekLast() T {
	return q.Queue[len(q.Queue)-1]
}

func (q *Queue[T]) Length() int {
	return len(q.Queue)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Length() == 0
}

func (q *Queue[T]) Push(value T) {
	q.Queue = append(q.Queue, value)
}

func (q *Queue[T]) Pop() T {
	first := q.Queue[len(q.Queue)-1]
	q.Queue = q.Queue[:len(q.Queue)-1]

	return first
}

func (q *Queue[T]) PopLeft() T {
	last := q.Queue[0]
	q.Queue = q.Queue[1:]

	return last
}

func (q *Queue[T]) Copy(node *Queue[T]) {
	q.Queue = node.Queue
}

func (q *Queue[T]) Clear() {
	q.Queue = make([]T, 0)
}

func (q *Queue[T]) Front() T {
	return q.Queue[0]
}

func (q *Queue[T]) Back() T {
	return q.Queue[len(q.Queue)-1]
}
