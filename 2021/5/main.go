package main

type Grid struct {
	squares [][]int
}

type Pos struct {
	x int
	y int
}

func NewGrid(h, w int) *Grid {
	cols := [][]int{}
	for i := 0; i < h; i++ {
		cols = append(cols, make([]int, w))
	}
	return &Grid{
		squares: cols,
	}
}

func (g *Grid) AddPoint(pos Pos) {
	g.squares[pos.y][pos.x] += 1
}

func createLinePoints(from, to Pos) []Pos {
	points := []Pos{}
	if from.x == to.x {
		// increment y's
		for y := from.y; y < to.y; y++ {
			points = append(points, Pos{
				x: from.x,
				y: y,
			})
		}
	} else if from.y == to.y {
		// increment x's
		for x := from.x; x < to.x; x++ {
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
	for _, point := range linePoints {
		g.AddPoint(point)
	}
}

type Line struct {
	from Pos
	to   Pos
}

func NewGridFromLines(lines []Line) *Grid {
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
	return NewGrid(maxY, maxX)
}

func main() {}
