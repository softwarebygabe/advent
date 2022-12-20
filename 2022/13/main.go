package main

import (
	"fmt"
	"strings"

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

func composeLists(s string) [][]int {
	lists := [][]int{}
	currItems := []int{}
	// stillStart := true
	for i, r := range s {
		if r == '[' {
			// if stillStart {
			// 	continue
			// }
			if len(currItems) > 0 {
				lists = append(lists, currItems)
				currItems = []int{}
			}
			continue
		}
		if r == ',' {
			continue
		}
		if r == ']' {
			if i == len(s)-1 {
				break
			}
			if len(currItems) > 0 {
				lists = append(lists, currItems)
				currItems = []int{}
			}
			continue
		}
		// stillStart = false
		currItems = append(currItems, util.MustParseInt(string(r)))
	}
	if len(currItems) > 0 {
		lists = append(lists, currItems)
	}
	return lists
}

func inCorrectOrder(l, r [][]int) bool {
	for i, leftItems := range l {
		if len(r)-1 < i {
			// right ran out of items
			return false
		}
		rightItems := r[i]
		minLen := len(leftItems)
		if len(rightItems) < minLen {
			minLen = len(rightItems)
		}
		for j := 0; j < minLen; j++ {
			leftItem := leftItems[j]
			rightItem := rightItems[j]
			if rightItem < leftItem {
				return false
			}
		}
	}
	return true
}

func composeAndCompare(ls, rs string) bool {
	leftQueue := createQueue(ls)
	rightQueue := createQueue(rs)
	for !leftQueue.Empty() {
		if rightQueue.Empty() {
			return false
		}
		leftList, _ := leftQueue.Dequeue()
		rightList, _ := rightQueue.Dequeue()
		if !compare(leftList, rightList) {
			return false
		}
	}
	return true
}

func createQueue(s string) *util.Queue[[]int] {
	elems := strings.Split(s, ",")
	fmt.Println(elems)
	return util.NewQueue[[]int]()
}

func compare(l, r []int) bool {
	minLen := len(l)
	if len(r) < minLen {
		minLen = len(r)
	}
	for i := 0; i < minLen; i++ {
		leftItem := l[i]
		rightItem := r[i]
		if rightItem < leftItem {
			return false
		}
	}
	return true
}

func part1() {
	pairs := parseInput("input_test.txt")
	correctIndices := []int{}
	for idx, pair := range pairs {
		// fmt.Println(pair)
		left := pair[0]
		right := pair[1]
		leftItemsLists := composeLists(left)
		rightItemsLists := composeLists(right)
		fmt.Println(left)
		fmt.Println(right)
		fmt.Println(leftItemsLists)
		fmt.Println(rightItemsLists)
		if inCorrectOrder(leftItemsLists, rightItemsLists) {
			fmt.Println("in the right order: true")
			correctIndices = append(correctIndices, idx+1)
		}
		fmt.Println("====")
	}
	fmt.Println(correctIndices)
}

func part1Take2() {
	pairs := parseInput("input_test.txt")
	correctIndices := []int{}
	for i, pair := range pairs {
		left := pair[0]
		right := pair[1]
		fmt.Println(left)
		fmt.Println(right)
		if composeAndCompare(left, right) {
			correctIndices = append(correctIndices, i+1)
		}
		fmt.Println("====")
	}
	fmt.Println(correctIndices)
}

func main() {
	part1Take2()
}
