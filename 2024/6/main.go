package main

import (
	"fmt"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type Square struct {
	value    string
	visited  bool
	visitDir []util.Direction
}

func (s *Square) isGuard() bool {
	return s.value == "^"
}

func (s *Square) isObstacle() bool {
	return s.value == "#"
}

func (s *Square) visit() {
	s.visited = true
}

func (s *Square) beenHere(dir util.Direction) bool {
	if s.visited {
		for _, d := range s.visitDir {
			if d == dir {
				return true
			}
		}
	}
	return false
}

func (s *Square) visitWithDir(dir util.Direction) {
	s.visited = true
	s.visitDir = append(s.visitDir, dir)
}

func parseToGrid(filename string) (util.Grid[*Square], error) {
	lines, err := util.ReadInput(filename, util.ReaderToStrings)
	if err != nil {
		return nil, err
	}
	grid := util.NewGrid[*Square]()
	for _, line := range lines {
		chars := strings.Split(line, "")
		sqs := []*Square{}
		for _, char := range chars {
			sqs = append(sqs, &Square{
				value:    char,
				visitDir: make([]util.Direction, 0),
			})
		}
		grid = append(grid, sqs)
	}
	return grid, nil
}

func markVisited(grid util.Grid[*Square], pos util.Position) {
	sq, ok := grid.Get(pos)
	if !ok {
		panic("markVisted: invalid pos")
	}
	sq.visit()
}

func markVisitedDir(grid util.Grid[*Square], pos util.Position, dir util.Direction) {
	sq, ok := grid.Get(pos)
	if !ok {
		panic("markVisted: invalid pos")
	}
	sq.visitWithDir(dir)
}

func markObstacle(grid util.Grid[*Square], pos util.Position) {
	sq, ok := grid.Get(pos)
	if !ok {
		panic("markObstacle: invalid pos")
	}
	sq.value = "#"
}

func Part2(filename string) {
	grid, err := parseToGrid(filename)
	if err != nil {
		panic(err)
	}
	loopPositions := []util.Position{}

	// markObstacle(grid, util.NewPosition(7, 6))
	// fmt.Println("looped?", moveGuardThroughGrid(grid))

	grid.ForEach(func(p util.Position, v *Square) {
		newGrid, _ := parseToGrid(filename)
		if !v.isObstacle() {
			fmt.Println("checking p:", p.String())
			markObstacle(newGrid, p)
			looped := moveGuardThroughGrid(newGrid)
			fmt.Println("looped?", looped)
			if looped {
				loopPositions = append(loopPositions, p)
			}
		}
	})
	// for _, p := range loopPositions {
	// 	fmt.Println(p.String())
	// }
	fmt.Println("result:", len(loopPositions))
}

func moveGuardThroughGrid(grid util.Grid[*Square]) bool {
	// find the current position of the guard
	var currGuardPos util.Position
	grid.ForEach(func(p util.Position, v *Square) {
		if v.isGuard() {
			currGuardPos = p
		}
	})
	// move the guard through through it's path until it leaves the grid
	currGuardDir := util.Up
	for {
		currSq, _ := grid.Get(currGuardPos)
		been := currSq.beenHere(currGuardDir)
		// mark the current place and dir as visited
		markVisitedDir(grid, currGuardPos, currGuardDir)
		// fmt.Println("currPos:", currGuardPos.String(), "currDir:", currGuardDir, "currSq:", currSq, "been here?", been)
		if been {
			fmt.Println("loop detected")
			return true
		}
		// check the next move
		nextSq, ok := grid.Get(currGuardPos.Move(currGuardDir, 1))
		if !ok {
			// next move is out of bounds
			return false
		}
		if nextSq.isObstacle() {
			// if the next move is into an obstacle, turn
			newDir := currGuardDir.Turn(util.Right)
			currGuardDir = newDir
		} else {
			// next move is not an obstacle, move forward
			newPos := currGuardPos.Move(currGuardDir, 1)
			currGuardPos = newPos
		}
	}
}

func Part1(filename string) {
	grid, err := parseToGrid(filename)
	if err != nil {
		panic(err)
	}

	// find the current position of the guard
	var currGuardPos util.Position
	grid.ForEach(func(p util.Position, v *Square) {
		if v.isGuard() {
			currGuardPos = p
		}
	})
	// move the guard through through it's path until it leaves the grid
	markVisited(grid, currGuardPos)
	currGuardDir := util.Up
	for {
		// move
		nextPos := currGuardPos.Move(currGuardDir, 1)
		nextSq, ok := grid.Get(nextPos)
		if !ok {
			// we have left the grid, end
			markVisited(grid, currGuardPos)
			break
		}
		if nextSq.isObstacle() {
			// turn guard 90deg right
			currGuardDir = currGuardDir.Turn(util.Right)
		} else {
			// move forward
			markVisited(grid, currGuardPos)
			currGuardPos = nextPos
		}
	}

	// count all the visited squares
	var visitedSum int
	grid.ForEach(func(p util.Position, v *Square) {
		if v.visited {
			visitedSum++
		}
	})
	// fmt.Println(grid)
	// fmt.Println(currGuardPos)
	fmt.Println("result:", visitedSum)
}

func main() {
	// Part1("input_ex.txt")
	// Part1("input_1.txt")
	// Part2("input_ex.txt")
	Part2("input_1.txt")
}
