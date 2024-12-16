package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/softwarebygabe/advent/pkg/colors"
	"github.com/softwarebygabe/advent/pkg/util"
	"gonum.org/v1/gonum/stat/combin"
)

type Node struct {
	value      string
	isAntinode bool
}

func (n Node) Antenna() bool {
	return n.value != "."
}

func (n Node) Antinode() bool {
	return n.isAntinode
}

func (n *Node) MarkAntinode() {
	n.isAntinode = true
}

func (n Node) String() string {
	if n.isAntinode {
		if n.value == "." {
			return colors.Sprintf(colors.Yellow, "%s", "#")
		}
		return colors.Sprintf(colors.Yellow, "%s", n.value)
	}
	return n.value
}

func parseGrid(filename string) util.Grid[*Node] {
	lines, err := util.ReadInput(filename, util.ReaderToStrings)
	if err != nil {
		panic(err)
	}
	grid := util.NewGrid[*Node]()
	for _, line := range lines {
		nodes := []*Node{}
		for _, char := range strings.Split(line, "") {
			nodes = append(nodes, &Node{
				value: char,
			})
		}
		grid = append(grid, nodes)
	}
	return grid
}

func maybeMarkAntinodes(grid util.Grid[*Node], pos ...util.Position) {
	for _, p := range pos {
		n, ok := grid.Get(p)
		if ok {
			n.MarkAntinode()
		}
	}
}

func markAntinode(grid util.Grid[*Node], pos util.Position) error {
	n, ok := grid.Get(pos)
	if !ok {
		return errors.New("invalid position")
	}
	n.MarkAntinode()
	return nil
}

func Part1(filename string) {
	grid := parseGrid(filename)
	// create a map of frequencies to antenna positions
	antennaPositions := make(map[string][]util.Position)
	grid.ForEach(func(p util.Position, v *Node) {
		if v.Antenna() {
			positions, ok := antennaPositions[v.value]
			if !ok {
				antennaPositions[v.value] = []util.Position{p}
			} else {
				antennaPositions[v.value] = append(positions, p)
			}
		}
	})
	// for each uniq antenna frequency, create all combinations of positions and compute antinode positions
	for _, v := range antennaPositions {
		// fmt.Println(k, v)
		combins := combin.Combinations(len(v), 2) // all combinations of pairs of positions
		for _, c := range combins {
			// fmt.Println(i, c)
			p1, p2 := v[c[0]], v[c[1]]
			dy := p2.Row - p1.Row
			dx := p2.Col - p1.Col
			// get the delta positions from p1 and p2
			// note: doing the delta from one of the two will equal the other (p2),
			// in that case invert dy and dx and do again for that point (p2)
			dp1 := p2.Delta(dy, dx)
			dp2 := p1.Delta(-dy, -dx)
			maybeMarkAntinodes(grid, dp1, dp2)
		}
	}
	grid.Print(os.Stdout)
	// count up the antinodes
	var antinodeCount int
	grid.ForEach(func(p util.Position, v *Node) {
		if v.Antinode() {
			antinodeCount++
		}
	})
	fmt.Println("result:", antinodeCount)
}

func Part2(filename string) {
	grid := parseGrid(filename)
	// create a map of frequencies to antenna positions
	antennaPositions := make(map[string][]util.Position)
	grid.ForEach(func(p util.Position, v *Node) {
		if v.Antenna() {
			positions, ok := antennaPositions[v.value]
			if !ok {
				antennaPositions[v.value] = []util.Position{p}
			} else {
				antennaPositions[v.value] = append(positions, p)
			}
		}
	})
	// for each uniq antenna frequency, create all combinations of positions and compute antinode positions
	for _, v := range antennaPositions {
		// fmt.Println(k, v)
		combins := combin.Combinations(len(v), 2) // all combinations of pairs of positions
		for _, c := range combins {
			// fmt.Println(i, c)
			p1, p2 := v[c[0]], v[c[1]]
			dy := p2.Row - p1.Row
			dx := p2.Col - p1.Col
			maybeMarkAntinodes(grid, p1, p2)
			// get the delta positions from p1 and p2
			// note: doing the delta from one of the two will equal the other (p2),
			// in that case invert dy and dx and do again for that point (p2)
			factor := 1
			for {
				dp := p2.Delta(dy*factor, dx*factor)
				if err := markAntinode(grid, dp); err != nil {
					break
				}
				factor++
			}
			factor = 1
			for {
				dp := p1.Delta(-dy*factor, -dx*factor)
				if err := markAntinode(grid, dp); err != nil {
					break
				}
				factor++
			}
		}
	}
	grid.Print(os.Stdout)
	// count up the antinodes
	var antinodeCount int
	grid.ForEach(func(p util.Position, v *Node) {
		if v.Antinode() {
			antinodeCount++
		}
	})
	fmt.Println("result:", antinodeCount)
}

func main() {
	// Part1("input_ex.txt")
	// Part1("input_1.txt")
	Part2("input_ex.txt")
	Part2("input_1.txt")
}
