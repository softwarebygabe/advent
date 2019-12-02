package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type evaluator = func(line string)

func doForAllLines(filepath string, eval evaluator) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		eval(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func calculateFuel(mass float64) int {
	possibleFuel := int(mass/3) - 2
	if possibleFuel < 0 {
		return 0
	}
	return possibleFuel
}

func main() {

	sum := 0

	recalcQueue := []int{}

	eval := func(i string) {
		// convert to a float
		f, err := strconv.ParseFloat(i, 64)
		if err != nil {
			log.Fatal(err)
		}

		moreFuel := calculateFuel(f)
		sum += moreFuel
		recalcQueue = append(recalcQueue, moreFuel)

	}

	doForAllLines("./input.txt", eval)

	// now go through the recalcQueue
	for len(recalcQueue) > 0 {
		last := recalcQueue[len(recalcQueue)-1]
		moreFuel := calculateFuel(float64(last))
		if moreFuel > 0 {
			recalcQueue[len(recalcQueue)-1] = moreFuel
			sum += moreFuel
		} else {
			recalcQueue = recalcQueue[0 : len(recalcQueue)-1]
		}
	}

	fmt.Println(sum)
}
