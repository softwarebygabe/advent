package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/softwarebygabe/advent/pkg/colors"
	"github.com/softwarebygabe/advent/pkg/util"
)

type Octopus struct {
	energy  int
	flashes int
	flashed bool
}

func (o *Octopus) excite() {
	o.energy++
	if o.energy > 9 {
		o.energy = 0
		o.flashes++
		o.flashed = true
	}
}

func parseInput(f string) Grid {
	grid := [][]*Octopus{}
	util.EvalEachLine(f, func(line string) {
		row := []*Octopus{}
		for _, intS := range strings.Split(line, "") {
			i, err := strconv.Atoi(intS)
			if err != nil {
				panic(err)
			}
			row = append(row, &Octopus{energy: i})
		}
		grid = append(grid, row)
	})
	return grid
}

func (g Grid) print() {
	for _, row := range g {
		for _, o := range row {
			if o.energy > 0 {
				fmt.Printf("%d", o.energy)
			} else {
				colors.Printf(colors.Yellow, "%d", o.energy)
			}
		}
		fmt.Printf("\n")
	}
}

type Grid [][]*Octopus

func (g Grid) computeStep() {
	// compute new energy levels
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			g.maybeBloom(y, x)
		}
	}
}

func (g Grid) maybeBloom(y, x int) {
	o := g[y][x]
	if !o.flashed {
		o.excite()
		if o.flashed {
			g.bloom(y, x)
		}
	}
}

func (g Grid) bloom(y, x int) {
	if y > 0 {
		g.maybeBloom(y-1, x)
		if x > 0 {
			g.maybeBloom(y-1, x-1)
		}
		if x < len(g[y])-1 {
			g.maybeBloom(y-1, x+1)
		}
	}
	if x > 0 {
		g.maybeBloom(y, x-1)
	}
	if x < len(g[y])-1 {
		g.maybeBloom(y, x+1)
	}
	if y < len(g)-1 {
		g.maybeBloom(y+1, x)
		if x > 0 {
			g.maybeBloom(y+1, x-1)
		}
		if x < len(g[y])-1 {
			g.maybeBloom(y+1, x+1)
		}
	}
}

func (g Grid) runSteps(n int) {
	curr := 1
	for curr <= n {
		g.computeStep()
		fmt.Println("After step", curr, ":")
		g.print()
		fmt.Println()
		// reset flashed
		for y := 0; y < len(g); y++ {
			for x := 0; x < len(g[y]); x++ {
				g[y][x].flashed = false
			}
		}
		curr++
	}
}

func (g Grid) runStepsUntilAllFlashing(n int) {
	curr := 1
	for curr <= n {
		g.computeStep()
		fmt.Println("After step", curr, ":")
		g.print()
		fmt.Println()
		if g.allFlashed() {
			fmt.Println("all flashing!")
			break
		}
		// reset flashed
		for y := 0; y < len(g); y++ {
			for x := 0; x < len(g[y]); x++ {
				g[y][x].flashed = false
			}
		}
		curr++
	}
}

func (g Grid) allFlashed() bool {
	all := true
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			if !g[y][x].flashed {
				all = false
				break
			}
		}
	}
	return all
}

func (g Grid) flashCount() int {
	var total int
	for _, row := range g {
		for _, o := range row {
			total += o.flashes
		}
	}
	return total
}

func Part1() {
	g := parseInput("./input.txt")
	g.print()
	fmt.Println()
	g.runSteps(100)
	fmt.Println("Total Flash Count:", g.flashCount())
}

func main() {
	g := parseInput("./input.txt")
	g.print()
	fmt.Println()
	g.runStepsUntilAllFlashing(1000)
}
