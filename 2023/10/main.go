package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

var directionsMap = map[string][]util.Direction{
	"|": {util.Up, util.Down},
	"-": {util.Left, util.Right},
	"L": {util.Up, util.Right},
	"J": {util.Up, util.Left},
	"7": {util.Left, util.Down},
	"F": {util.Right, util.Down},
}

type node struct {
	r string
	p util.Position
}

func newNode(r string, p util.Position) node {
	return node{r, p}
}

func (n node) String() string {
	return fmt.Sprintf("r=%s %s", n.r, n.p)
}

func stringString(lines []string) [][]string {
	res := [][]string{}
	for _, row := range lines {
		rowChars := strings.Split(row, "")
		res = append(res, rowChars)
	}
	return res
}

func part1(filename string) []node {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := stringString(util.ReaderToStrings(f))
	var start node
	var curr node
	var currInitialDir util.Direction
	var found bool
	// find the start
	for rowI, row := range lines {
		for colI, r := range row {
			if r == "S" {
				start = newNode(r, util.NewPosition(rowI, colI))
				// find curr to start on
				// look up
				if rowI > 0 {
					currR := lines[rowI-1][colI]
					switch currR {
					case "|", "7", "F":
						fmt.Println("found up")
						currInitialDir = util.Up
						curr = newNode(currR, util.NewPosition(rowI-1, colI))
					}
				}
				// look right
				if colI < len(row)-1 {
					currR := lines[rowI][colI+1]
					switch currR {
					case "-", "J", "7":
						fmt.Println("found right")
						currInitialDir = util.Right
						curr = newNode(currR, util.NewPosition(rowI, colI+1))
					}
				}
				// look down
				if rowI < len(lines)-1 {
					currR := lines[rowI+1][colI]
					switch currR {
					case "|", "L", "J":
						fmt.Println("found down")
						currInitialDir = util.Down
						curr = newNode(currR, util.NewPosition(rowI+1, colI))
					}
				}
				// look left
				if colI > 0 {
					currR := lines[rowI][colI-1]
					switch currR {
					case "-", "L", "F":
						fmt.Println("found left")
						currInitialDir = util.Left
						curr = newNode(currR, util.NewPosition(rowI, colI-1))
					}
				}
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	// find where to start

	// walk around
	loop := []node{curr}
	for curr.String() != start.String() {
		// fmt.Println(loop)
		back := currInitialDir.Opposite()
		dirs, ok := directionsMap[curr.r]
		if !ok {
			panic("ahh")
		}
		for _, dir := range dirs {
			if dir != back {
				// go this way
				newP := curr.p.Move(dir, 1)
				for i, r := range lines[newP.Row] {
					if i == newP.Col {
						n := newNode(r, newP)
						loop = append(loop, n)
						curr = n
						currInitialDir = dir
						break
					}
				}
			}
		}
	}
	// for _, n := range loop {
	// 	fmt.Println(n)
	// }
	fmt.Println(math.Floor(float64(len(loop) / 2)))
	return loop
}

func part2(filename string) {
	loop := part1(filename)
	// TODO: scan through and do subtraction on the col indexes
	// also do it on the row indexes and the diffs should be the areas
	// needed
}

func main() {
	part1("input.txt")
}
