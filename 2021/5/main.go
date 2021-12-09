package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type Grid struct {
	squares [][]int
}

type Pos struct {
	x int
	y int
}

func NewGrid(h, w int) *Grid {
	cols := [][]int{}
	for i := 0; i <= h; i++ {
		cols = append(cols, make([]int, w+1))
	}
	return &Grid{
		squares: cols,
	}
}

func (g Grid) String() string {
	result := ""
	for _, row := range g.squares {
		for _, count := range row {
			if count == 0 {
				result += ". "
			} else {
				result += fmt.Sprintf("%d ", count)
			}
		}
		result += "\n"
	}
	return result
}

func (g *Grid) AddPoint(pos Pos) {
	g.squares[pos.y][pos.x]++
}

func createSeries(from, to int) []int {
	result := []int{}
	if from < to {
		curr := from
		for curr <= to {
			result = append(result, curr)
			curr++
		}
		return result
	} else if from > to {
		curr := from
		for curr >= to {
			result = append(result, curr)
			curr--
		}
	}
	return result
}

func createDiagonals(from, to Pos) []Pos {
	points := []Pos{}
	// diagonal
	xseries := createSeries(from.x, to.x)
	yseries := createSeries(from.y, to.y)
	for i := 0; i < len(xseries); i++ {
		points = append(points, Pos{
			x: xseries[i],
			y: yseries[i],
		})
	}
	return points
}

func createLinePoints(from, to Pos) []Pos {
	points := []Pos{}
	if from.x == to.x {
		// increment y's
		for y := from.y; y <= to.y; y++ {
			points = append(points, Pos{
				x: from.x,
				y: y,
			})
		}
	} else if from.y == to.y {
		// increment x's
		for x := from.x; x <= to.x; x++ {
			points = append(points, Pos{
				x: x,
				y: from.y,
			})
		}
	}
	return points
}

func (g *Grid) AddLine(l Line) {
	linePoints := createLinePoints(l.from, l.to)
	if len(linePoints) < 1 {
		linePoints = createLinePoints(l.to, l.from)
	}
	if len(linePoints) < 1 {
		linePoints = createDiagonals(l.from, l.to)
	}
	// fmt.Println(linePoints)
	for _, point := range linePoints {
		g.AddPoint(point)
	}
}

type Line struct {
	from Pos
	to   Pos
}

func NewGridFromLines(lines []Line) *Grid {
	fmt.Println("finding max X and max Y...")
	// find the h and w
	var maxY, maxX int
	for _, line := range lines {
		// y
		if line.from.y > maxY {
			maxY = line.from.y
		}
		if line.to.y > maxY {
			maxY = line.to.y
		}
		// x
		if line.from.x > maxX {
			maxX = line.from.x
		}
		if line.to.x > maxX {
			maxX = line.to.x
		}
	}
	fmt.Println("found maxes")
	grid := NewGrid(maxY, maxX)
	fmt.Println("new empty grid created")
	// fmt.Println(lines)
	for _, line := range lines {
		grid.AddLine(line)
	}
	fmt.Println("lines added to grid")
	return grid
}

func parseInput(filename string) []Line {
	fmt.Println("parsing input...")
	lines := []Line{}
	util.EvalEachLine(filename, func(line string) {
		pointStrings := strings.Split(line, " -> ")
		var newLine Line
		for i, ps := range pointStrings {
			splitList := strings.Split(ps, ",")
			// parse into Pos
			x, err := strconv.Atoi(splitList[0])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(splitList[1])
			if err != nil {
				panic(err)
			}
			newPos := Pos{x: x, y: y}
			if i == 0 {
				newLine.from = newPos
			} else {
				newLine.to = newPos
			}
		}
		lines = append(lines, newLine)
	})
	fmt.Println("input parsed")
	return lines
}

func (g Grid) CountIntersections() int {
	fmt.Println("counting intersections...")
	var intersectionCount int
	for _, row := range g.squares {
		for _, count := range row {
			if count > 1 {
				intersectionCount++
			}
		}
	}
	return intersectionCount
}

func main() {
	lines := parseInput("./input.txt")
	grid := NewGridFromLines(lines)
	// fmt.Println(grid.String())
	fmt.Println(grid.CountIntersections())
}
