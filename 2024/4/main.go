package main

import (
	"fmt"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

func parseToGrid(fname string) (util.Grid[string], error) {
	lines, err := util.ReadInput(fname, util.ReaderToStrings)
	if err != nil {
		return nil, err
	}
	grid := util.NewGrid[string]()
	for _, line := range lines {
		chars := strings.Split(line, "")
		grid = append(grid, chars)
	}
	return grid, nil
}

func Part1(fname string) {
	grid, err := parseToGrid(fname)
	if err != nil {
		panic(err)
	}

	xmasSum := 0
	grid.ForEach(func(p util.Position, v string) {
		if v == "X" {
			// look for XMAS in every direction
			searchForXMAS := func(dir util.Direction, start util.Position) bool {
				cursor := start
				for _, char := range []string{"M", "A", "S"} {
					cursor = cursor.Move(dir, 1)
					v2, ok := grid.Get(cursor)
					if !ok {
						return false
					}
					if v2 != char {
						return false
					}
					// else we found the next valid char!
				}
				// we found the whole XMAS in this dir!
				return true
			}
			for _, dir := range util.DirectionsAll {
				if searchForXMAS(dir, p) {
					xmasSum++
				}
			}
		}
	})
	fmt.Println("result:", xmasSum)

}

func Part2(fname string) {
	grid, err := parseToGrid(fname)
	if err != nil {
		panic(err)
	}

	sum := 0
	grid.ForEach(func(p util.Position, v string) {
		if v == "A" {
			// A is at the center of the X so look at the spokes
			// two Ms and two Ss need to be on opposite corners
			corners := []util.Position{
				p.Move(util.UpLeft, 1),
				p.Move(util.UpRight, 1),
				p.Move(util.DownRight, 1),
				p.Move(util.DownLeft, 1),
			}
			validCornerConfigs := []string{
				"MSSM", "SSMM", "SMMS", "MMSS",
			}
			cornerConfig := ""
			for _, corner := range corners {
				v, ok := grid.Get(corner)
				if !ok {
					return
				}
				cornerConfig += v
			}
			valid := false
			for _, validConfig := range validCornerConfigs {
				if cornerConfig == validConfig {
					valid = true
					break
				}
			}
			if valid {
				sum++
			}
		}
	})
	fmt.Println("result:", sum)
}

func main() {
	// Part1("input_ex.txt")
	// Part1("input_1.txt")
	Part2("input_ex.txt")
	Part2("input_1.txt")
}
