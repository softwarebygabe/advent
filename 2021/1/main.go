package main

import (
	"fmt"
	"strconv"

	"github.com/softwarebygabe/advent/pkg/util"
)

func parseInput(filename string) []int {
	depths := []int{}
	util.EvalEachLine(filename, func(line string) {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		depths = append(depths, i)
	})
	return depths
}

func countIncreasing(depths []int) int {
	var numIncreasing int
	for i, depth := range depths {
		if i > 0 {
			prev := depths[i-1]
			if prev < depth {
				numIncreasing++
			}
		}
	}
	return numIncreasing
}

func countIncreasingWindows(depths []int, windowSize int) int {
	windowSums := []int{}
	var windowStart int
	var windowEnd int
	for i := 0; i < len(depths); i++ {
		windowStart = i
		windowEnd = i + (windowSize - 1)
		if windowEnd >= len(depths) {
			break
		}
		currI := windowStart
		localSum := 0
		for currI <= windowEnd {
			depth := depths[currI]
			localSum += depth
			currI++
		}
		windowSums = append(windowSums, localSum)
	}
	// fmt.Println(windowSums)
	return countIncreasing(windowSums)
}

func main() {
	depths := parseInput("./input.txt")
	fmt.Println(countIncreasingWindows(depths, 3))
}
