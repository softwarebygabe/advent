package util

import "fmt"

type Queue[T any] []T

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	*q = append(*q, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(*q) == 0 {
		var v T
		return v, false
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v, true
}

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) String() string {
	return fmt.Sprintf("%v", *q)
}
