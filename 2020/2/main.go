package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type lineEvaluator = func(line string)

func doForAllInputLines(filepath string, fn lineEvaluator) {
	// open file
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// scan through file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		// parse an input line
		line := scanner.Text()
		fn(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func mustParseLine(line string) (min, max int, letter, password string) {
	splitOnColon := strings.Split(line, ": ")
	password = splitOnColon[1]

	splitOnSpace := strings.Split(splitOnColon[0], " ")
	letter = splitOnSpace[1]

	splitOnHyphen := strings.Split(splitOnSpace[0], "-")
	min, err := strconv.Atoi(splitOnHyphen[0])
	if err != nil {
		log.Fatal(err)
	}
	max, err = strconv.Atoi(splitOnHyphen[1])
	if err != nil {
		log.Fatal(err)
	}
	return
}

func isValid(min, max int, letter, password string) bool {
	fmt.Println(min, max, letter, password)
	var letterCount int
	pLetters := strings.Split(password, "")
	for _, pLetter := range pLetters {
		if pLetter == letter {
			letterCount++
		}
	}
	return min <= letterCount && letterCount <= max
}

func RunTest() {
	var validCount int
	var lineEval lineEvaluator = func(line string) {
		min, max, letter, password := mustParseLine(line)
		if isValid(min, max, letter, password) {
			validCount++
		}
	}
	doForAllInputLines("input_test.txt", lineEval)
	fmt.Println(validCount)
}

func RunPart1() {
	var validCount int
	var lineEval lineEvaluator = func(line string) {
		min, max, letter, password := mustParseLine(line)
		if isValid(min, max, letter, password) {
			validCount++
		}
	}
	doForAllInputLines("input.txt", lineEval)
	fmt.Println(validCount)
}

func RunPart2() {
	var validCount int
	var evalLine lineEvaluator = func(line string) {
		pos1, pos2, letter, password := mustParseLine(line)

		var isValid bool
		pwdLetters := strings.Split(password, "")

		inPos1 := pwdLetters[pos1-1] == letter
		inPos2 := pwdLetters[pos2-1] == letter

		if inPos1 && !inPos2 {
			isValid = true
		}
		if inPos2 && !inPos1 {
			isValid = true
		}

		if isValid {
			validCount++
		}

	}
	doForAllInputLines("input.txt", evalLine)
	fmt.Println(validCount)
}

func main() {
	RunPart2()
}
