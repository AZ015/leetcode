package stack

import "strings"

type Stack struct {
	stack []string
}

func New() *Stack {
	return &Stack{
		stack: make([]string, 0),
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	}

	return s.stack[len(s.stack)-1]
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	last := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return last
}

func (s *Stack) ToString(sep string) string {
	if s.IsEmpty() {
		return ""
	}

	return strings.Join(s.stack, sep)
}

func (s *Stack) Push(el string) {
	s.stack = append(s.stack, el)
}
