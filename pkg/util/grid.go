package util

import (
	"fmt"
	"io"
)

type Grid[T any] [][]T

func NewGrid[T any]() Grid[T] {
	return Grid[T]{}
}

func (g Grid[T]) Empty() bool {
	return len(g) < 1
}

func (g Grid[T]) Valid(p Position) bool {
	return !g.Empty() && p.Row >= 0 && p.Row < len(g) && p.Col >= 0 && p.Col < len(g[p.Row])
}

func (g Grid[T]) Get(p Position) (T, bool) {
	if !g.Valid(p) {
		var v T
		return v, false
	}
	return g[p.Row][p.Col], true
}

func (g Grid[T]) Print(w io.Writer) {
	for _, row := range g {
		for _, v := range row {
			fmt.Fprintf(w, "%v", v)
		}
		fmt.Fprintf(w, "\n")
	}
}

func (g Grid[T]) ForEach(f func(p Position, v T)) {
	for y, r := range g {
		for x, v := range r {
			f(NewPosition(y, x), v)
		}
	}
}
