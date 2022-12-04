package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type sectionAssignment struct {
	from int
	to   int
}

func (s sectionAssignment) contains(s2 sectionAssignment) bool {
	return s.from <= s2.from && s2.to <= s.to
}

func (s sectionAssignment) overlaps(s2 sectionAssignment) bool {
	return (s2.from <= s.to && s.from <= s2.to) ||
		(s.from <= s2.to && s2.from <= s.to)
}

func parseInput(filepath string) [][]sectionAssignment {
	results := [][]sectionAssignment{}
	util.EvalEachLine(filepath, func(line string) {
		pairs := strings.Split(line, ",")
		pairAssignment := []sectionAssignment{}
		for _, pair := range pairs {
			nums := strings.Split(pair, "-")
			from, err := strconv.Atoi(nums[0])
			if err != nil {
				panic(err)
			}
			to, err := strconv.Atoi(nums[1])
			if err != nil {
				panic(err)
			}
			pairAssignment = append(pairAssignment, sectionAssignment{
				from: from,
				to:   to,
			})
		}
		results = append(results, pairAssignment)
	})
	return results
}

func part1() {
	sectionAssignments := parseInput("input.txt")
	var countOverlapping int
	for _, pairAssignment := range sectionAssignments {
		if pairAssignment[0].contains(pairAssignment[1]) || pairAssignment[1].contains(pairAssignment[0]) {
			countOverlapping++
		}
	}
	fmt.Println("num overlapping:", countOverlapping)
}

func part2() {
	sectionAssignments := parseInput("input.txt")
	var countOverlapping int
	for _, pairAssignment := range sectionAssignments {
		a1 := pairAssignment[0]
		a2 := pairAssignment[1]
		if a1.overlaps(a2) {
			countOverlapping++
		}
	}
	fmt.Println("num overlapping:", countOverlapping)
}

func main() {
	// part1()
	part2()
}
