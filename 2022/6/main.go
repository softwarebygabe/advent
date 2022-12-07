package main

import (
	"fmt"

	"github.com/softwarebygabe/advent/pkg/util"
)

func parseInput(filepath string) string {
	lines := []string{}
	util.EvalEachLine(filepath, func(line string) {
		lines = append(lines, line)
	})
	return lines[0]
}

func containsDuplicateChars(s string) bool {
	seen := map[rune]struct{}{}
	for _, r := range s {
		_, inMap := seen[r]
		if !inMap {
			seen[r] = struct{}{}
		} else {
			return true
		}
	}
	return false
}

func part1() {
	data := parseInput("input.txt")
	if len(data) < 4 {
		panic("datastream invalid")
	}
	var result int
	// slide through with window 4-char wide
	for i := 4; i < len(data); i++ {
		subString := data[i-4 : i]
		fmt.Println(subString)
		if !containsDuplicateChars(subString) {
			result = i
			break
		}
	}
	fmt.Println("result:", result)
}

func part2() {
	data := parseInput("input_test.txt")
	if len(data) < 4 {
		panic("datastream invalid")
	}
	var result int
	// slide through with window 14-char wide
	for i := 14; i < len(data); i++ {
		subString := data[i-14 : i]
		fmt.Println(subString)
		if !containsDuplicateChars(subString) {
			result = i
			break
		}
	}
	fmt.Println("result:", result)
}

func main() {
	part2()
}
