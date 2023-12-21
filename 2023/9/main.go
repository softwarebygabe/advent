package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type history struct {
	readings []int
}

func newHistory(readings ...int) history {
	h := history{
		readings: make([]int, 0),
	}
	h.readings = readings
	return h
}

func (h history) predictNext() int {
	fmt.Println("predictNext")
	// create diffs
	fullSeries := [][]int{h.readings}
	currSeries := h.readings
	for {
		diffs := make([]int, len(currSeries)-1)
		if len(diffs) < 1 {
			break
		}
		for i := 0; i < len(currSeries)-1; i++ {
			diffs[i] = currSeries[i+1] - currSeries[i]
		}
		// check if all zeros
		allZeros := true
		for _, n := range diffs {
			if n != 0 {
				allZeros = false
				break
			}
		}
		// else add diffs to the readings
		fullSeries = append(fullSeries, diffs)
		currSeries = diffs
		if allZeros {
			// we are done with diffs
			break
		}
	}
	fmt.Println("fullSeries:", fullSeries)
	// go backwards up through and calc next
	for i := len(fullSeries) - 1; i >= 0; i-- {
		currSeries := fullSeries[i]
		var nextVal int
		if i != len(fullSeries)-1 {
			prevSeries := fullSeries[i+1]
			// add up the last val of the curr series and last val of the prev series
			nextVal = currSeries[len(currSeries)-1] + prevSeries[len(prevSeries)-1]
		}
		// add it in
		fullSeries[i] = append(currSeries, nextVal)
	}
	fmt.Println("fullSeries:", fullSeries)
	firstSeries := fullSeries[0]
	return firstSeries[len(firstSeries)-1]
}

func (h history) predictNext2() int {
	fmt.Println("predictNext")
	// create diffs
	fullSeries := [][]int{h.readings}
	currSeries := h.readings
	for {
		diffs := make([]int, len(currSeries)-1)
		if len(diffs) < 1 {
			break
		}
		for i := 0; i < len(currSeries)-1; i++ {
			diffs[i] = currSeries[i+1] - currSeries[i]
		}
		// check if all zeros
		allZeros := true
		for _, n := range diffs {
			if n != 0 {
				allZeros = false
				break
			}
		}
		// else add diffs to the readings
		fullSeries = append(fullSeries, diffs)
		currSeries = diffs
		if allZeros {
			// we are done with diffs
			break
		}
	}
	fmt.Println("fullSeries:", fullSeries)
	// go backwards up through and calc next
	for i := len(fullSeries) - 1; i >= 0; i-- {
		currSeries := fullSeries[i]
		var nextVal int
		if i != len(fullSeries)-1 {
			prevSeries := fullSeries[i+1]
			// add up the last val of the curr series and last val of the prev series
			nextVal = currSeries[0] - prevSeries[0]
		}
		// add it in
		newSeries := []int{nextVal}
		for _, n := range currSeries {
			newSeries = append(newSeries, n)
		}
		fullSeries[i] = newSeries
	}
	fmt.Println("fullSeries:", fullSeries)
	firstSeries := fullSeries[0]
	return firstSeries[0]
}

func part1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)
	histories := []history{}
	for _, line := range lines {
		numsRaw := strings.Split(line, " ")
		nums := []int{}
		for _, nRaw := range numsRaw {
			n, err := strconv.Atoi(nRaw)
			if err != nil {
				panic(err)
			}
			nums = append(nums, n)
		}
		histories = append(histories, newHistory(nums...))
	}
	predictions := []int{}
	for _, h := range histories {
		predictions = append(predictions, h.predictNext())
	}
	fmt.Println("sum:", util.Sum(predictions...))
}

func part2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)
	histories := []history{}
	for _, line := range lines {
		numsRaw := strings.Split(line, " ")
		nums := []int{}
		for _, nRaw := range numsRaw {
			n, err := strconv.Atoi(nRaw)
			if err != nil {
				panic(err)
			}
			nums = append(nums, n)
		}
		histories = append(histories, newHistory(nums...))
	}
	predictions := []int{}
	for _, h := range histories {
		predictions = append(predictions, h.predictNext2())
	}
	fmt.Println("sum:", util.Sum(predictions...))
}

func main() {
	part2("input.txt")
}
