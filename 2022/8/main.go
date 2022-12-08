package main

import (
	"fmt"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

func parseInput(filepath string) [][]int {
	results := [][]int{}
	util.EvalEachLine(filepath, func(line string) {
		nums := strings.Split(line, "")
		currLine := []int{}
		for _, n := range nums {
			currLine = append(currLine, util.MustParseInt(n))
		}
		results = append(results, currLine)
	})
	return results
}

func isTreeVisible(grid [][]int, x, y int) bool {
	if y == 0 || y == len(grid)-1 {
		return true // outside top/bottom edges
	}
	if x == 0 || x == len(grid[y])-1 {
		return true // outside sides
	}
	tree := grid[y][x]
	// look along row
	row := grid[y]
	leftClear := true
	rightClear := true
	for xi, ti := range row {
		if xi < x {
			if ti >= tree && leftClear {
				leftClear = false
			}
		}
		if x < xi {
			if ti >= tree && rightClear {
				rightClear = false
			}
		}
	}
	// look top/bottom
	topClear := true
	bottomClear := true
	for yi := 0; yi < len(grid); yi++ {
		ti := grid[yi][x]
		if yi < y {
			// top
			if ti >= tree && topClear {
				topClear = false
			}
		}
		if y < yi {
			if ti >= tree && bottomClear {
				bottomClear = false
			}
		}
	}
	return leftClear || rightClear || topClear || bottomClear
}

func calcScenicScore(grid [][]int, x, y int) int {
	tree := grid[y][x]
	var leftScore int
	if x != 0 {
		for xi := x - 1; xi >= 0; xi-- {
			ti := grid[y][xi]
			leftScore++
			if ti >= tree {
				break
			}
		}
	}
	var rightScore int
	if x != len(grid[y]) {
		for xi := x + 1; xi < len(grid[y]); xi++ {
			ti := grid[y][xi]
			rightScore++
			if ti >= tree {
				break
			}
		}
	}
	// look top/bottom
	var topScore int
	if y != 0 {
		for yi := y - 1; yi >= 0; yi-- {
			ti := grid[yi][x]
			topScore++
			if ti >= tree {
				break
			}
		}
	}
	var bottomScore int
	if y != len(grid) {
		for yi := y + 1; yi < len(grid); yi++ {
			ti := grid[yi][x]
			bottomScore++
			if ti >= tree {
				break
			}
		}
	}
	score := leftScore * rightScore * topScore * bottomScore
	fmt.Println("tree:", tree, "score:", score, "l:", leftScore, "r:", rightScore, "t:", topScore, "b:", bottomScore)
	return score
}

func part1() {
	trees := parseInput("input.txt")
	fmt.Println(trees)
	// go through all trees and say if it's vis
	visTreeCount := 0
	for y, row := range trees {
		for x, _ := range row {
			vis := isTreeVisible(trees, x, y)
			if vis {
				visTreeCount++
			}
		}
	}
	fmt.Println("visible trees:", visTreeCount)
}

func part2() {
	trees := parseInput("input.txt")
	fmt.Println(trees)
	// go through all trees and say if it's vis
	maxScore := 0
	for y, row := range trees {
		for x, _ := range row {
			score := calcScenicScore(trees, x, y)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	fmt.Println("max score:", maxScore)
}

func main() {
	part2()
}
