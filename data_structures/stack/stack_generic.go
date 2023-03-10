package stack

type StackGeneric[T any] struct {
	stack []T
}

func NewStackGeneric[T any]() *StackGeneric[T] {
	return &StackGeneric[T]{
		stack: make([]T, 0),
	}
}

func (s *StackGeneric[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *StackGeneric[T]) Peek() T {
	return s.stack[len(s.stack)-1]
}

func (s *StackGeneric[T]) Pop() T {
	last := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return last
}

func (s *StackGeneric[T]) Push(el T) {
	s.stack = append(s.stack, el)
}
