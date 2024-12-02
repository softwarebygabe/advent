package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

// returns true if Safe, false if Unsafe
func safetyEval(report []int, min, max int) bool {
	var increasing bool
	for i, v := range report {
		if i+1 == len(report) {
			break // we have reached the last item
		}
		v2 := report[i+1]

		if v == v2 {
			return false
		}

		if v < v2 {
			// we have an increase
			if i != 0 && !increasing {
				// we have an inconsistency
				return false
			}
			increasing = true
		} else {
			// we have a decrease
			if i != 0 && increasing {
				// we have an inconsistency
				return false
			}
			increasing = false
		}

		delta := int(math.Abs(float64(v) - float64(v2)))
		if delta < min || delta > max {
			return false
		}
	}
	return true
}

func removeIndex(v []int, idx int) []int {
	res := []int{}
	for i, vi := range v {
		if i != idx {
			res = append(res, vi)
		}
	}
	return res
}

// returns true if Safe, false if Unsafe
func safetyEvalWithDampener(report []int, min, max int) bool {
	if !safetyEval(report, min, max) {
		for i := range report {
			if safetyEval(removeIndex(report, i), min, max) {
				return true
			}
		}
		return false
	}
	return true
}

func Part1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)
	safeCount := 0
	for _, line := range lines {
		lineStrings := strings.Split(line, " ")
		lineInts := util.StringsToInts(lineStrings)
		res := safetyEval(lineInts, 1, 3)
		fmt.Println("eval:", res)
		if res {
			safeCount++
		}
	}
	fmt.Println("safe count:", safeCount)
}

func Part2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)
	safeCount := 0
	for _, line := range lines {
		lineStrings := strings.Split(line, " ")
		lineInts := util.StringsToInts(lineStrings)
		res := safetyEvalWithDampener(lineInts, 1, 3)
		fmt.Println("eval:", res)
		if res {
			safeCount++
		}
	}
	fmt.Println("safe count:", safeCount)
}

func main() {
	// Part1("input_part1.txt")
	Part2("input_part1.txt")
}
