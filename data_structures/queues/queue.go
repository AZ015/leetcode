package queues

type Queue[T any] struct {
	queue []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		queue: make([]T, 0),
	}
}

func (q *Queue[T]) Length() int {
	return len(q.queue)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Length() == 0
}

func (q *Queue[T]) Push(value T) {
	q.queue = append(q.queue, value)
}

func (q *Queue[T]) Pop() T {
	first := q.queue[0]
	q.queue = q.queue[1:]

	return first
}

func (q *Queue[T]) Peek() T {
	return q.queue[0]
}
