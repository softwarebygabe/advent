package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type Entry struct {
	signalPatterns []string
	output         []string
}

func parseInput(filename string) []Entry {
	entries := []Entry{}
	util.EvalEachLine(filename, func(line string) {
		splitList := strings.Split(line, " | ")
		entries = append(entries, Entry{
			signalPatterns: strings.Split(splitList[0], " "),
			output:         strings.Split(splitList[1], " "),
		})
	})
	return entries
}

func digit(segment string) (int, error) {
	switch len(segment) {
	case 2:
		return 1, nil
	case 4:
		return 4, nil
	case 3:
		return 7, nil
	case 7:
		return 8, nil
	}
	return 0, errors.New("unable to find digit")
}

func Part1() {
	entries := parseInput("input.txt")
	digitCount := map[int]int{
		1: 0,
		4: 0,
		7: 0,
		8: 0,
	}
	for _, entry := range entries {
		for _, output := range entry.output {
			digit, err := digit(output)
			if err != nil {
				continue
			}
			digitCount[digit] += 1
		}
	}
	// sum digits
	var sum int
	for _, v := range digitCount {
		sum += v
	}
	fmt.Println(sum)
}

func includes(s, sub []string) bool {
	founds := []bool{}
	for _, a := range sub {
		for _, b := range s {
			if a == b {
				founds = append(founds, true)
			}
		}
	}
	for _, f := range founds {
		if !f {
			return false
		}
	}
	return true
}

func diff(a, b []string) []string {
	diff := map[string]int{}
	for _, ai := range a {
		diff[ai] = 0
	}
	for _, bi := range b {
		_, ok := diff[bi]
		if !ok {
			diff[bi] = 0
		} else {
			delete(diff, bi)
		}
	}
	results := []string{}
	for k := range diff {
		results = append(results, k)
	}
	return results
}

var numberToSegmentMap = map[int][]int{
	0: {0, 1, 2, 4, 5, 6},    // done
	6: {0, 1, 3, 4, 5, 6},    // done
	9: {0, 1, 2, 3, 5, 6},    // done
	2: {0, 2, 3, 4, 6},       // done
	3: {0, 2, 3, 5, 6},       // done
	5: {0, 1, 3, 5, 6},       // done
	1: {2, 5},                // done
	7: {0, 2, 5},             // done
	4: {1, 2, 3, 5},          // done
	8: {0, 1, 2, 3, 4, 5, 6}, // done
}

var letters = []string{"a", "b", "c", "d", "e", "f", "g"}

type Key = map[int][]string

func patternToDigit(key map[int][]string, pattern string) (int, error) {
	patternLetters := strings.Split(pattern, "")
	for digit, letters := range key {
		// check len
		if len(letters) == len(patternLetters) && len(diff(letters, patternLetters)) == 0 {
			return digit, nil
		}
	}
	return 0, errors.New("no digit found in key")
}

func createKey(entry Entry) map[int][]string {
	key := make(map[int][]string)
	// add in the easy patterns first
	for _, pattern := range entry.signalPatterns {
		digit, err := digit(pattern)
		if err != nil {
			continue
		}
		key[digit] = strings.Split(pattern, "")
	}
	// now try the hard patterns

	// letter 0 is the one in 7 but not in 1
	l0 := diff(key[7], key[1])[0]
	// fmt.Println(l0)
	// 9 is the letters in 4 + l0 + l6
	l6 := ""
	for _, pattern := range entry.signalPatterns {
		patternLetters := strings.Split(pattern, "")
		fourPlusL0 := append(key[4], l0)
		if len(patternLetters) > len(fourPlusL0) {
			diff := diff(patternLetters, fourPlusL0)
			if len(diff) == 1 {
				key[9] = patternLetters
				l6 = diff[0]
			}
		}
	}
	// fmt.Println(l6)
	// l4 is the diff of 4 + l0 + l6 and 8
	l4 := diff(key[8], append(key[4], l0, l6))[0]
	// fmt.Println(l4)
	// ls13 is the diff of 1 and 4
	ls13 := diff(key[1], key[4])
	// fmt.Println(ls13)
	// l5 is l0 + ls13 + l6 + 5
	l5 := ""
	for _, pattern := range entry.signalPatterns {
		patternLetters := strings.Split(pattern, "")
		looker := append(ls13, l0, l6)
		if len(patternLetters) > len(looker) {
			diff := diff(patternLetters, looker)
			if len(diff) == 1 {
				key[5] = patternLetters
				l5 = diff[0]
			}
		}
	}
	// fmt.Println(l5)
	// 6 is l0 + ls13 + l4 + l5 + l6
	key[6] = append(ls13, l0, l4, l5, l6)
	// l2 is diff of 9 and l0 + ls13 + l5 + l6
	l2 := diff(key[9], append(ls13, l0, l5, l6))[0]
	l1 := ""
	for _, pattern := range entry.signalPatterns {
		patternLetters := strings.Split(pattern, "")
		looker := []string{l0, l2, l4, l5, l6}
		if len(patternLetters) > len(looker) {
			diff := diff(patternLetters, looker)
			if len(diff) == 1 {
				key[0] = patternLetters
				l1 = diff[0]
			}
		}
	}
	l3 := diff(ls13, []string{l1})[0]
	key[2] = []string{l0, l2, l3, l4, l6}
	key[3] = []string{l0, l2, l3, l5, l6}
	// fmt.Println(ls13)
	// fmt.Println(l1)
	// fmt.Println(l3)
	return key
}

func main() {
	entries := parseInput("input.txt")
	var sum int
	for _, entry := range entries {
		key := createKey(entry)
		// fmt.Println(key)
		var number int
		fmt.Println(entry.output)
		for i, pattern := range entry.output {
			digit, err := patternToDigit(key, pattern)
			if err != nil {
				panic(err)
			}
			switch i {
			case 0:
				number += digit * 1000
			case 1:
				number += digit * 100
			case 2:
				number += digit * 10
			default:
				number += digit
			}
		}
		// fmt.Println(number)
		sum += number
	}
	fmt.Println(sum)
}
