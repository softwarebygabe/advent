package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// given a list of entries (int) find the two
// that add up to 2020 and multiply them together
func part1(entries []int) int {
	for i, entryI := range entries {
		for j, entryJ := range entries {
			if i != j {
				// if they are not the same entry then add
				sum := entryI + entryJ
				if sum == 2020 {
					return entryI * entryJ
				}
			}
		}
	}
	return 0
}

func part2(entries []int) int {
	for i, ei := range entries {
		for j, ej := range entries {
			for k, ek := range entries {
				if i != j && i != k && j != k {
					sum := ei + ej + ek
					if sum == 2020 {
						return ei * ej * ek
					}
				}
			}
		}
	}
	return 0
}

func mustParseInput(filepath string) []int {
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

func RunTestPart1() {
	testInput := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	output := part1(testInput)
	fmt.Println(output)
}

func RunPart1() {
	part1Input := mustParseInput("./input.txt")
	output := part1(part1Input)
	fmt.Println(output)
}

func RunTestPart2() {
	testInput := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	output := part2(testInput)
	fmt.Println(output)
}

func RunPart2() {
	part2Input := mustParseInput("./input.txt")
	output := part2(part2Input)
	fmt.Println(output)
}

func main() {
	RunPart2()
}
