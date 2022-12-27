package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/softwarebygabe/advent/pkg/util"
)

func parseInput(filepath string) [][]string {
	pairs := [][]string{}
	currPair := []string{}
	util.EvalEachLine(filepath, func(line string) {
		if line == "" {
			pairs = append(pairs, currPair)
			currPair = []string{}
		} else {
			currPair = append(currPair, line)
		}
	})
	return append(pairs, currPair)
}

type iterator struct {
	raw     string
	values  []value
	currIdx int
}

func newIterator(s string) *iterator {
	iter := &iterator{raw: s, currIdx: -1}
	// parse to values
	values := []value{}
	var opencount, closercount int
	runes := []rune(s)
	var currValueString string
	for i := 1; i < len(runes)-1; i++ {
		char := fmt.Sprintf("%c", runes[i])
		switch char {
		case "[":
			// if opencount > 0 {
			// 	currValueString += char
			// }
			opencount++
			currValueString += char
		case "]":
			if opencount > 0 {
				currValueString += char
			}
			closercount++
			if closercount == opencount {
				// reset
				if currValueString != "" {
					values = append(values, value(currValueString))
					currValueString = ""
				}
				opencount = 0
				closercount = 0
			}
		case ",":
			if opencount > 0 {
				currValueString += char
			} else {
				// reset
				if currValueString != "" {
					values = append(values, value(currValueString))
					currValueString = ""
				}
				opencount = 0
				closercount = 0
			}
		default:
			currValueString += char
		}
	}
	// reset
	if currValueString != "" {
		values = append(values, value(currValueString))
		currValueString = ""
	}
	iter.values = values
	return iter
}

func (i *iterator) next() bool {
	i.currIdx++
	return i.currIdx < len(i.values)
}

func (i *iterator) value() value {
	var idx int
	if i.currIdx > -1 {
		idx = i.currIdx
	}
	return i.values[idx]
}

func (i *iterator) reset() {
	i.currIdx = -1
}

type value string

func (v value) isArray() bool {
	if len(v) < 1 {
		return false
	}
	return []rune(v)[0] == '['
}

func (v value) int() int {
	res, _ := strconv.Atoi(string(v))
	return res
}

func _calculate(l, r *iterator) (result bool, shorted bool) {
	for l.next() {
		if !r.next() {
			// right ran out of items
			return false, true
		}
		lv := l.value()
		rv := r.value()
		fmt.Println("- Compare", lv, "vs", rv)
		switch {
		case !lv.isArray() && !rv.isArray():
			if lv.int() < rv.int() {
				// left side is smaller
				fmt.Println("left side is smaller so correct order")
				return true, true
			}
			if lv.int() > rv.int() {
				// right side is smaller
				fmt.Println("right side is smaller so NOT in right order")
				return false, true
			}
		case lv.isArray() && rv.isArray():
			li := newIterator(string(lv))
			ri := newIterator(string(rv))
			res, ok := _calculate(li, ri)
			if ok {
				return res, true
			}
		default:
			// mixed types
			fmt.Println("mixed types")
			ls := string(lv)
			rs := string(rv)
			if !lv.isArray() {
				ls = "[" + ls + "]"
			}
			if !rv.isArray() {
				rs = "[" + rs + "]"
			}
			li := newIterator(string(ls))
			ri := newIterator(string(rs))
			res, ok := _calculate(li, ri)
			if ok {
				return res, true
			}
		}
	}
	// left ran out of items
	if !r.next() {
		// right also ran out of items
		fmt.Println("reached the end of current left and right")
		return false, false
	}
	fmt.Println("left ran out of items so in the correct order")
	return true, true
}

func part1() {
	pairs := parseInput("input.txt")
	correctIndices := []int{}
	for idx, pair := range pairs {
		// fmt.Println(pair)
		left := pair[0]
		right := pair[1]
		fmt.Println("left:", left)
		fmt.Println("right:", right)
		leftIter := newIterator(left)
		rightIter := newIterator(right)
		inCorrectOrder, _ := _calculate(leftIter, rightIter)
		fmt.Println("inCorrectOrder:", inCorrectOrder)
		if inCorrectOrder {
			correctIndices = append(correctIndices, idx+1)
		}
		fmt.Println("====")
	}
	fmt.Println(correctIndices)
	var sum int
	for _, i := range correctIndices {
		sum += i
	}
	fmt.Println("result:", sum)
}

type ByPacketOrder []*iterator

func (a ByPacketOrder) Len() int      { return len(a) }
func (a ByPacketOrder) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPacketOrder) Less(i, j int) bool {
	a[i].reset()
	a[j].reset()
	inOrder, _ := _calculate(a[i], a[j])
	return inOrder
}

func part2() {
	pairs := parseInput("input.txt")
	// turn them into iterators
	allIterators := []*iterator{}
	for _, pair := range pairs {
		for _, packet := range pair {
			allIterators = append(allIterators, newIterator(packet))
		}
	}
	// add in the divider packets
	dividers := []string{"[[2]]", "[[6]]"}
	for _, d := range dividers {
		allIterators = append(allIterators, newIterator(d))
	}
	// sort the list of iterators
	sort.Sort(ByPacketOrder(allIterators))
	// list them out
	packetsInOrder := []string{}
	for _, iter := range allIterators {
		fmt.Println(iter.raw)
		packetsInOrder = append(packetsInOrder, iter.raw)
	}
	mult := 1
	for i, p := range packetsInOrder {
		if p == dividers[0] || p == dividers[1] {
			mult *= (i + 1)
		}
	}
	fmt.Println("result:", mult)
}

func main() {
	part2()
}
