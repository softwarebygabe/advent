package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

func parseInput(filepath string) [][]point {
	results := [][]point{}
	util.EvalEachLine(filepath, func(line string) {
		pointStrings := strings.Split(line, "->")
		linePoints := []point{}
		for _, ps := range pointStrings {
			xy := strings.Split(strings.Trim(ps, " "), ",")
			p := point{
				x: util.MustParseInt(xy[0]),
				y: util.MustParseInt(xy[1]),
			}
			linePoints = append(linePoints, p)
		}
		results = append(results, linePoints)
	})
	return results
}

type point struct {
	x, y int
}

func newPoint(x, y int) point {
	return point{x: x, y: y}
}

func keyFromPoint(p point) string {
	return key(p.x, p.y)
}

func key(x, y int) string {
	return fmt.Sprintf("(%d,%d)", x, y)
}

type cave struct {
	rockMap    map[string]struct{}
	sandMap    map[string]struct{}
	sandSource point
	maxRockY   int
	floorY     int
}

func newCave(rocks []point) *cave {
	c := &cave{
		rockMap:    make(map[string]struct{}),
		sandMap:    make(map[string]struct{}),
		sandSource: newPoint(500, 0),
	}
	for _, rock := range rocks {
		if c.maxRockY < rock.y {
			c.maxRockY = rock.y
		}
		c.rockMap[keyFromPoint(rock)] = struct{}{}
	}
	// for part2
	c.floorY = c.maxRockY + 2
	return c
}

func (c *cave) isRock(x, y int) bool {
	_, inMap := c.rockMap[key(x, y)]
	return inMap
}

func (c *cave) isSand(x, y int) bool {
	_, inMap := c.sandMap[key(x, y)]
	return inMap
}

func (c *cave) isAir(x, y int) bool {
	return !c.isRock(x, y) && !c.isSand(x, y) && y < c.floorY
}

func (c *cave) placeSand(p point) {
	c.sandMap[keyFromPoint(p)] = struct{}{}
}

func (c *cave) dropSand() error {
	pos := newPoint(c.sandSource.x, c.sandSource.y)
	for pos.y <= c.maxRockY {
		if pos.y == c.maxRockY {
			return errors.New("overflow")
		}
		// try down
		if c.isAir(pos.x, pos.y+1) {
			pos.y++
			continue
		}
		// try down-left
		if c.isAir(pos.x-1, pos.y+1) {
			pos.x--
			pos.y++
			continue
		}
		// try down-right
		if c.isAir(pos.x+1, pos.y+1) {
			pos.x++
			pos.y++
			continue
		}
		// if all three blocked, place sand
		c.placeSand(pos)
		break
	}
	return nil
}

func (c *cave) dropSandWithFloor() error {
	pos := newPoint(c.sandSource.x, c.sandSource.y)
	for {
		// try down
		if c.isAir(pos.x, pos.y+1) {
			pos.y++
			continue
		}
		// try down-left
		if c.isAir(pos.x-1, pos.y+1) {
			pos.x--
			pos.y++
			continue
		}
		// try down-right
		if c.isAir(pos.x+1, pos.y+1) {
			pos.x++
			pos.y++
			continue
		}
		// if all three blocked, place sand, unless we haven't moved
		if pos.x == c.sandSource.x && pos.y == c.sandSource.y {
			return errors.New("source covered")
		}
		c.placeSand(pos)
		break
	}
	return nil
}

func calculateLinePoints(p1, p2 point) []point {
	line := []point{}
	if p1.x == p2.x {
		minY := p1.y
		maxY := p1.y
		if p2.y < minY {
			minY = p2.y
		}
		if maxY < p2.y {
			maxY = p2.y
		}
		for i := minY; i <= maxY; i++ {
			line = append(line, newPoint(p1.x, i))
		}
	}
	if p1.y == p2.y {
		minX := p1.x
		maxX := p1.x
		if p2.x < minX {
			minX = p2.x
		}
		if maxX < p2.x {
			maxX = p2.x
		}
		for i := minX; i <= maxX; i++ {
			line = append(line, newPoint(i, p1.y))
		}
	}
	return line
}

func part1() {
	allRockLinePoints := parseInput("input.txt")
	allRockPoints := []point{}
	for _, rockLine := range allRockLinePoints {
		for i := 1; i < len(rockLine); i++ {
			rockPoints := calculateLinePoints(rockLine[i-1], rockLine[i])
			allRockPoints = append(allRockPoints, rockPoints...)
		}
	}
	fmt.Println(allRockPoints)
	cave := newCave(allRockPoints)
	var sandGrains int
	var err error
	for err == nil {
		err = cave.dropSand()
		if err == nil {
			sandGrains++
		}
	}
	fmt.Println(err)
	fmt.Println("sand grains:", sandGrains)
}

func part2() {
	allRockLinePoints := parseInput("input.txt")
	allRockPoints := []point{}
	for _, rockLine := range allRockLinePoints {
		for i := 1; i < len(rockLine); i++ {
			rockPoints := calculateLinePoints(rockLine[i-1], rockLine[i])
			allRockPoints = append(allRockPoints, rockPoints...)
		}
	}
	fmt.Println(allRockPoints)
	cave := newCave(allRockPoints)
	var sandGrains int
	var err error
	for err == nil {
		err = cave.dropSandWithFloor()
		sandGrains++
	}
	fmt.Println(err)
	fmt.Println("sand grains:", sandGrains)
}

func main() {
	part2()
}
