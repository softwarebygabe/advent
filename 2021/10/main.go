package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

var pointMap = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var completionPointMap = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

var openToClose = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var closeToOpen = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

func isOpener(s string) bool {
	_, inMap := openToClose[s]
	return inMap
}

func isCloser(s string) bool {
	_, inMap := closeToOpen[s]
	return inMap
}

func pop(l []string) ([]string, string) {
	if len(l) > 0 {
		v := l[len(l)-1]
		return l[:len(l)-1], v
	}
	return l, ""
}

func push(l []string, s string) []string {
	return append(l, s)
}

func findIllegalChar(line string) string {
	var illegal string
	openStack := []string{}
	for _, s := range strings.Split(line, "") {
		if isOpener(s) {
			openStack = push(openStack, s)
		}
		if isCloser(s) {
			mostRecentOpener := openStack[len(openStack)-1]
			if mostRecentOpener != closeToOpen[s] {
				// illegal closing char, does not match most recent opener
				illegal = s
				break
			} else {
				// valid close of chunk, pop the most recent opener
				openStack, _ = pop(openStack)
			}
		}
	}
	return illegal
}

func completeLine(line string) string {
	var completion string
	openStack := []string{}
	for _, s := range strings.Split(line, "") {
		if isOpener(s) {
			openStack = push(openStack, s)
		}
		if isCloser(s) {
			openStack, _ = pop(openStack)
		}
	}
	// now close out the open stack working back
	for len(openStack) > 0 {
		_, opener := pop(openStack)
		completion += openToClose[opener]
		openStack, _ = pop(openStack)
	}
	return completion
}

func Part1() {
	illegalChars := []string{}
	badLines := []string{}
	util.EvalEachLine("./input.txt", func(line string) {
		illegalChar := findIllegalChar(line)
		if illegalChar != "" {
			badLines = append(badLines, line)
			illegalChars = append(illegalChars, illegalChar)
		}
	})
	// print results
	for i := 0; i < len(badLines); i++ {
		fmt.Println(badLines[i], "illegal char:", illegalChars[i])
	}
	// calculate score
	var total int
	for _, s := range illegalChars {
		total += pointMap[s]
	}
	fmt.Println("Total Score:", total)
}

func Part2() {
	// filter out corrupted lines
	incompleteLines := []string{}
	util.EvalEachLine("./input.txt", func(line string) {
		illegalChar := findIllegalChar(line)
		if illegalChar == "" {
			incompleteLines = append(incompleteLines, line)
		}
	})
	// compute completions
	lineCompletions := []string{}
	for _, l := range incompleteLines {
		completion := completeLine(l)
		lineCompletions = append(lineCompletions, completion)
		fmt.Println(l, "completion:", completion)
	}
	// compute line scores
	lineTotals := []int{}
	for _, completion := range lineCompletions {
		var lineTotal int
		for _, s := range strings.Split(completion, "") {
			lineTotal = lineTotal * 5
			lineTotal += completionPointMap[s]
		}
		lineTotals = append(lineTotals, lineTotal)
		fmt.Println(completion, "-", lineTotal, "total points.")
	}
	// find middle score
	sort.Ints(lineTotals)
	fmt.Println("len:", len(lineTotals), "half:", len(lineTotals)/2)
	fmt.Println(lineTotals[len(lineTotals)/2])
}

func main() {
	Part2()
}
