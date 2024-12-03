package util

import "fmt"

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(*s) == 0 {
		var v T
		return v, false
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(*s) == 0 {
		var v T
		return v, false
	}
	v := (*s)[len(*s)-1]
	return v, true
}

func (s *Stack[T]) Len() int {
	return len(*s)
}

// Reset empties all entries in the stack.
func (s *Stack[T]) Reset() {
	for s.Len() > 0 {
		s.Pop()
	}
}
