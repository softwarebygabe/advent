package main

import (
	"fmt"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

func scoreTrailhead(grid util.Grid[int], trailhead util.Position) (part1 int, part2 int) {
	dirs := []util.Direction{
		util.Up,
		util.Down,
		util.Left,
		util.Right,
	}
	var part2Score int
	trailEndMap := make(map[string]struct{})
	next := util.NewQueue[util.Position]()
	next.Enqueue(trailhead)
	for !next.Empty() {
		pos, ok := next.Dequeue()
		if !ok {
			panic("next queue is empty!")
		}
		curr, ok := grid.Get(pos)
		if !ok {
			panic("next pos is off grid!")
		}
		// look in all dirs
		for _, dir := range dirs {
			nextPos := pos.Move(dir, 1)
			peek, ok := grid.Get(nextPos)
			if ok {
				// on grid
				if peek-1 == curr {
					// increases by exactly 1
					if peek == 9 {
						trailEndMap[nextPos.String()] = struct{}{}
						part2Score++
					} else {
						next.Enqueue(nextPos) // keep walking
					}
				}
			}
		}
	}
	return len(trailEndMap), part2Score
}

func Part1(filename string) {
	lines, err := util.ReadInput(filename, util.ReaderToStrings)
	if err != nil {
		panic(err)
	}
	grid := util.NewGrid[int]()
	for _, line := range lines {
		row := []int{}
		for _, char := range strings.Split(line, "") {
			row = append(row, util.MustParseInt(char))
		}
		grid = append(grid, row)
	}

	// find trailheads
	trailheads := []util.Position{}
	grid.ForEach(func(p util.Position, v int) {
		if v == 0 {
			trailheads = append(trailheads, p)
		}
	})
	var sum1, sum2 int
	for _, th := range trailheads {
		part1, part2 := scoreTrailhead(grid, th)
		sum1 += part1
		sum2 += part2
	}
	fmt.Println("sum1", sum1)
	fmt.Println("sum2", sum2)
}

func main() {
	Part1("input_ex.txt")
	Part1("input_1.txt")
}
