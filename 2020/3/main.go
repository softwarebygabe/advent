package main

import (
	"fmt"
	"strings"

	"github.com/softwarebygabe/advent/pkg/helpers"
)

type Square struct {
	IsTree bool
}

type Map [][]Square

func parseMap(inputFile string) Map {
	result := make([][]Square, 0)
	helpers.EvalEachLine(inputFile, func(line string) {
		row := []Square{}
		splitList := strings.Split(line, "")
		for _, s := range splitList {
			if s == "." {
				row = append(row, Square{IsTree: false})
			} else if s == "#" {
				row = append(row, Square{IsTree: true})
			} else {
				panic("encountered an unknown map square")
			}
		}
		result = append(result, row)
	})
	return result
}

func generateBiggerMap(slopeMap Map, right, down int) Map {
	yLength := len(slopeMap)
	amountOfDowns := yLength / down
	desiredXLength := right * amountOfDowns
	xLength := len(slopeMap[0])
	repeatAmount := desiredXLength / xLength
	repeatAmount++ // safety
	// repeat each line pattern
	biggerMap := make([][]Square, 0)
	for _, row := range slopeMap {
		newRow := []Square{}
		var amt int
		for amt <= repeatAmount {
			newRow = append(newRow, row...)
			amt++
		}
		biggerMap = append(biggerMap, newRow)
	}
	return biggerMap
}

func countTrees(slopeMap Map, right, down int) int {
	var treeCount int

	var x int
	var y int

	var finished bool

	for !finished {
		// check current position for tree
		currSquare := slopeMap[y][x]
		if currSquare.IsTree {
			treeCount++
		}
		// check if new position is valid
		newY := y + down
		if newY < len(slopeMap) {
			newX := x + right
			if newX < len(slopeMap[newY]) {
				// if both new positions are valid move on
				y = newY
				x = newX
			} else {
				finished = true
			}
		} else {
			finished = true
		}
	}

	return treeCount
}

func main() {
	slopeMap := parseMap("./input.txt")

	slopePairs := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	results := []int{}
	for _, pair := range slopePairs {
		right := pair[0]
		down := pair[1]
		biggerMap := generateBiggerMap(slopeMap, right, down)
		treeCount := countTrees(biggerMap, right, down)
		results = append(results, treeCount)
	}
	fmt.Println(results)
	var mult int
	for i, r := range results {
		if i != 0 {
			mult = mult * r
		} else {
			mult = r
		}
	}
	fmt.Println(mult)
}
