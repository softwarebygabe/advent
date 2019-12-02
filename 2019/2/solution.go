package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func add(startIndex int, fullInput []int) ([]int, error) {
	arg1I := fullInput[startIndex+1]
	arg2I := fullInput[startIndex+2]
	outputI := fullInput[startIndex+3]

	if arg1I >= len(fullInput) || arg2I >= len(fullInput) || outputI >= len(fullInput) {
		return fullInput, errors.New("")
	}

	a := fullInput[arg1I]
	b := fullInput[arg2I]
	sum := a + b
	fullInput[outputI] = sum
	return fullInput, nil
}

func multiply(startIndex int, fullInput []int) ([]int, error) {
	arg1I := fullInput[startIndex+1]
	arg2I := fullInput[startIndex+2]
	outputI := fullInput[startIndex+3]

	if arg1I >= len(fullInput) || arg2I >= len(fullInput) || outputI >= len(fullInput) {
		return fullInput, errors.New("")
	}

	a := fullInput[arg1I]
	b := fullInput[arg2I]
	mult := a * b
	fullInput[outputI] = mult
	return fullInput, nil
}

func computer(intcode []int) ([]int, error) {
	cursor := 0
	stop := false
	for !stop {
		opcode := intcode[cursor]
		switch opcode {
		case 1:
			intcode, err := add(cursor, intcode)
			if err != nil {
				return intcode, err
			}
		case 2:
			intcode, err := multiply(cursor, intcode)
			if err != nil {
				return intcode, err
			}
		case 99:
			stop = true
		}
		newCursor := cursor + 4
		if newCursor >= len(intcode) {
			stop = true
		}
		cursor = newCursor
	}
	return intcode, nil
}

func parseInput(filepath string) []int {
	result := []int{}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numberStrings := strings.Split(line, ",")
		for _, num := range numberStrings {
			integer, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, int(integer))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func checkOutput(intcode []int, test int) bool {
	return intcode[0] == test
}

func main() {

	lookingFor := 19690720

	noun := 0
	for noun < 100 {
		verb := 0
		for verb < 100 {

			program := parseInput("./input.txt")

			program[1] = noun
			program[2] = verb

			output, _ := computer(program)

			if checkOutput(output, lookingFor) {
				fmt.Println(fmt.Sprintf("noun: %d verb: %d", noun, verb))
				fmt.Println(100*noun + verb)
				break
			}

			verb++
		}

		noun++
	}

}
