package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type knot struct {
	x, y int
}

func (k *knot) move(x, y int) {
	k.x = x
	k.y = y
}

func (k *knot) up() {
	k.move(k.x, k.y+1)
}

func (k *knot) down() {
	k.move(k.x, k.y-1)
}

func (k *knot) left() {
	k.move(k.x-1, k.y)
}

func (k *knot) right() {
	k.move(k.x+1, k.y)
}

func newKnot() *knot {
	return &knot{}
}

func parseInput(filepath string) []string {
	results := []string{}
	util.EvalEachLine(filepath, func(line string) {
		results = append(results, line)
	})
	return results
}

type board struct {
	headVisit map[string]int
	tailVisit map[string]int
}

func newBoard() *board {
	b := &board{
		headVisit: make(map[string]int),
		tailVisit: make(map[string]int),
	}
	b.visitHead(0, 0)
	b.visitTail(0, 0)
	return b
}

func (b *board) coordKey(x, y int) string {
	return fmt.Sprintf("(%d,%d)", x, y)
}

func (b *board) visitHead(x, y int) {
	b.headVisit[b.coordKey(x, y)]++
}

func (b *board) visitTail(x, y int) {
	b.tailVisit[b.coordKey(x, y)]++
}

func intAbs(v int) int {
	absF := math.Abs(float64(v))
	return int(absF)
}

func movePair(b *board, h, t *knot, dir string) (tailMoved bool) {
	// move h
	switch dir {
	case "U":
		h.up()
	case "D":
		h.down()
	case "R":
		h.right()
	case "L":
		h.left()
	default:
		panic("dir not supported")
	}
	fmt.Println("head:", b.coordKey(h.x, h.y))
	dx := h.x - t.x
	dy := h.y - t.y
	fmt.Println("dx", dx, "dy", dy)
	if 1 < intAbs(dx) || 1 < intAbs(dy) {
		// tail needs to move towards head
		if dx > 0 {
			t.right()
		}
		if dy > 0 {
			t.up()
		}
		if dx < 0 {
			t.left()
		}
		if dy < 0 {
			t.down()
		}
		fmt.Println("tail moved")
		tailMoved = true
	}
	fmt.Println("tail:", b.coordKey(t.x, t.y))
	return
}

func moveRope(b *board, rope []*knot, dir string) {
	var prevKnot *knot
	for _, knot := range rope {
		if prevKnot == nil {
			switch dir {
			case "U":
				knot.up()
			case "D":
				knot.down()
			case "R":
				knot.right()
			case "L":
				knot.left()
			default:
				panic("dir not supported")
			}
		} else {
			dx := prevKnot.x - knot.x
			dy := prevKnot.y - knot.y
			// fmt.Println("dx", dx, "dy", dy)
			if 1 < intAbs(dx) || 1 < intAbs(dy) {
				// tail needs to move towards head
				if dx > 0 {
					knot.right()
				}
				if dy > 0 {
					knot.up()
				}
				if dx < 0 {
					knot.left()
				}
				if dy < 0 {
					knot.down()
				}
			}
		}
		prevKnot = knot
	}
}

func runMove(b *board, h, t *knot, dir string, num int) {
	// fmt.Println("==", dir, num, "==")
	for i := 0; i < num; i++ {
		movePair(b, h, t, dir)
		// record visits
		b.visitHead(h.x, h.y)
		b.visitTail(t.x, t.y)
	}
}

func part1() {
	directionStrings := parseInput("input.txt")
	head, tail, board := newKnot(), newKnot(), newBoard()
	for _, directionString := range directionStrings {
		elems := strings.Split(directionString, " ")
		dir := elems[0]
		num := util.MustParseInt(elems[1])
		runMove(board, head, tail, dir, num)
	}
	fmt.Println("tailVisits len:", len(board.tailVisit))
}

func part2() {
	directionStrings := parseInput("input.txt")
	board := newBoard()
	rope := []*knot{}
	for i := 0; i < 10; i++ {
		rope = append(rope, newKnot())
	}
	for _, directionString := range directionStrings {
		elems := strings.Split(directionString, " ")
		dir := elems[0]
		num := util.MustParseInt(elems[1])
		// fmt.Println("==", dir, num, "==")
		for i := 0; i < num; i++ {
			// move all knots in the rope
			moveRope(board, rope, dir)
			// head := rope[0]
			tail := rope[len(rope)-1]
			// board.visitHead(head.x, head.y)
			board.visitTail(tail.x, tail.y)
		}
		// for _, knot := range rope {
		// 	fmt.Println(board.coordKey(knot.x, knot.y))
		// }
	}
	fmt.Println("tailVisits len:", len(board.tailVisit))
}

func main() {
	part2()
}
