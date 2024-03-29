package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/softwarebygabe/advent/pkg/util"
)

func parseInput(filename string) [][]int {
	results := [][]int{}
	currElf := []int{}
	util.EvalEachLine(filename, func(line string) {
		if line != "" {
			i, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			currElf = append(currElf, i)
		} else {
			results = append(results, currElf)
			currElf = []int{}
		}
	})
	if len(currElf) > 0 {
		results = append(results, currElf)
		currElf = []int{}
	}
	return results
}

func sum(l []int) int {
	var res int
	for _, v := range l {
		res += v
	}
	return res
}

func Part1() {
	elves := parseInput("./input.txt")
	var maxCals int
	for _, elf := range elves {
		elfCals := sum(elf)
		if maxCals < elfCals {
			maxCals = elfCals
		}
	}
	fmt.Println("max cals:", maxCals)
}

func Part2() {
	elves := parseInput("./input.txt")
	elfTotals := []int{}
	for _, elf := range elves {
		elfTotals = append(elfTotals, sum(elf))
	}
	// sort the elf totals
	sort.Ints(elfTotals)
	largestThree := elfTotals[len(elfTotals)-3:]
	fmt.Println(sum(largestThree))
}

func main() {
	Part2()
}
