package main

import (
	"fmt"
	"strconv"
	"strings"
)

// PuzzleInput ...
const PuzzleInput = "387638-919123"

// checks if the number given has 6 digits
func testDigitsRule(number string) bool {
	return len(strings.Split(number, "")) == 6
}

// checks if the number is within the puzzle input range
func testRangeRule(number string) bool {
	upperS := strings.Split(PuzzleInput, "-")[1]
	lowerS := strings.Split(PuzzleInput, "-")[0]

	upperI, _ := strconv.Atoi(upperS)
	lowerI, _ := strconv.Atoi(lowerS)

	numberI, _ := strconv.Atoi(number)

	return lowerI < numberI && numberI < upperI
}

// checks to make sure there are two adj numbers in number
func testAdjRule(number string) bool {
	digits := strings.Split(number, "")
	prev := ""
	adjCount := 0
	for _, digit := range digits {

		// no match check or move on
		if prev != digit {
			// we found a pair
			if adjCount == 1 {
				return true
			}
			// no pair move one
			prev = digit
			adjCount = 0
			continue
		}
		// if this matches prev
		if prev == digit {
			// if we have more than 1 adj move the loop
			if adjCount > 1 {
				continue
			}
			adjCount++
			continue
		}
	}
	// we reached the end check adj for end pair
	return adjCount == 1
}

// checks that all digits either increase or stay the same
func testNoDecreaseRule(number string) bool {
	digits := strings.Split(number, "")
	for i, digit := range digits {
		if i < len(digits)-1 {
			digitI, _ := strconv.Atoi(digit)
			nextDigitI, _ := strconv.Atoi(digits[i+1])
			if digitI > nextDigitI {
				return false
			}
		}
	}
	return true
}

func testAllRules(number int) bool {
	numberS := strconv.Itoa(number)
	return testAdjRule(numberS) && testRangeRule(numberS) && testDigitsRule(numberS) && testNoDecreaseRule(numberS)
}

func main() {
	fmt.Println("Hello world")

	upperS := strings.Split(PuzzleInput, "-")[1]
	lowerS := strings.Split(PuzzleInput, "-")[0]

	upperI, _ := strconv.Atoi(upperS)
	lowerI, _ := strconv.Atoi(lowerS)

	numMeetingCriteria := 0
	for test := lowerI + 1; test < upperI; test++ {
		if testAllRules(test) {
			numMeetingCriteria++
		}
	}

	fmt.Println(numMeetingCriteria)

	fmt.Println(testAdjRule("1111122"))
}
