package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type stack struct {
	crates []string
}

func newStack(crates ...string) *stack {
	return &stack{
		crates: crates,
	}
}

func (s *stack) String() string {
	result := ""
	for _, crate := range s.crates {
		result += crate
	}
	return result
}

func (s *stack) push(crate string) {
	s.crates = append(s.crates, crate)
}

func (s *stack) pop() string {
	if len(s.crates) < 1 {
		return ""
	}
	last := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]
	return last
}

func (s *stack) peek() string {
	if len(s.crates) < 1 {
		return ""
	}
	return s.crates[len(s.crates)-1]
}

type direction struct {
	move int
	from int
	to   int
}

func mustParseInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func parseInput(filepath string) ([]*stack, []direction) {
	directions := []direction{}
	stackLines := []string{}
	directionLines := []string{}
	var addToDir bool
	util.EvalEachLine(filepath, func(line string) {
		if line == "" {
			addToDir = true
			return
		}
		if addToDir {
			directionLines = append(directionLines, line)
		} else {
			stackLines = append(stackLines, line)
		}
	})
	for _, line := range directionLines {
		words := strings.Split(line, " ")
		if len(words) != 6 {
			panic("direction line does not fit format")
		}
		directions = append(directions, direction{
			move: mustParseInt(words[1]),
			from: mustParseInt(words[3]),
			to:   mustParseInt(words[5]),
		})
	}
	// go backwards through
	crateSets := [][]string{}
	for i := len(stackLines) - 2; 0 <= i; i-- {
		line := stackLines[i]
		crates := []string{}
		var rcount int
		var crate string
		for _, r := range line {
			switch {
			case rcount == 3:
				crates = append(crates, crate)
				rcount = 0
				crate = ""
			default:
				crate += string(r)
				rcount++
			}
		}
		crates = append(crates, crate)
		// fmt.Println(crates)
		// fmt.Println(len(crates))
		crateSets = append(crateSets, crates)
	}
	stacks := []*stack{}
	for range crateSets[0] {
		stacks = append(stacks, newStack())
	}
	for _, crateSet := range crateSets {
		for i, crate := range crateSet {
			val := string(crate[1])
			if val != " " {
				stacks[i].push(val)
			}
		}
	}
	return stacks, directions
}

func part1() {
	stacks, directions := parseInput("input.txt")
	fmt.Println(stacks)
	fmt.Println(directions)
	for _, dir := range directions {
		from := stacks[dir.from-1]
		to := stacks[dir.to-1]
		var moves int
		for moves < dir.move {
			to.push(from.pop())
			moves++
		}
	}
	var result string
	for _, s := range stacks {
		result += s.peek()
	}
	fmt.Println("result:", result)
}

func part2() {
	stacks, directions := parseInput("input.txt")
	for _, dir := range directions {
		from := stacks[dir.from-1]
		to := stacks[dir.to-1]
		switch dir.move {
		case 1:
			to.push(from.pop())
		default:
			var moves int
			tempStack := newStack()
			for moves < dir.move {
				tempStack.push(from.pop())
				moves++
			}
			for tempStack.peek() != "" {
				to.push(tempStack.pop())
			}
		}
	}
	var result string
	for _, s := range stacks {
		result += s.peek()
	}
	fmt.Println("result:", result)
}

func main() {
	part2()
}
