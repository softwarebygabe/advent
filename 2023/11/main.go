package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/softwarebygabe/advent/pkg/util"
)

func stringString(lines []string) [][]string {
	res := [][]string{}
	for _, row := range lines {
		rowChars := strings.Split(row, "")
		res = append(res, rowChars)
	}
	return res
}

func printImage(image [][]string) {
	for _, row := range image {
		fmt.Println(row)
	}
}

func expand(image [][]string) [][]string {
	isEmpty := func(set []string) bool {
		empty := true
		for _, v := range set {
			if v != "." {
				empty = false
				break
			}
		}
		return empty
	}
	// check all rows
	rowIdxs := []int{}
	for idx, row := range image {
		if isEmpty(row) {
			rowIdxs = append(rowIdxs, idx)
		}
	}
	// check all columns
	colIdxs := []int{}
	colDepth := len(image)
	for i := 0; i < colDepth; i++ {
		col := []string{}
		for _, row := range image {
			col = append(col, row[i])
		}
		if isEmpty(col) {
			colIdxs = append(colIdxs, i)
		}
	}
	fmt.Printf("rowIdxs=%v colIdxs=%v\n", rowIdxs, colIdxs)
	// insert blanks to expand
	includes := func(l []int, i int) bool {
		for _, v := range l {
			if v == i {
				return true
			}
		}
		return false
	}
	newImage := [][]string{}
	for idx, row := range image {
		// do cols first
		newRow := []string{}
		for cIdx, v := range row {
			if !includes(colIdxs, cIdx) {
				newRow = append(newRow, v)
			} else {
				// expand
				newRow = append(newRow, ".", ".")
			}
		}
		if !includes(rowIdxs, idx) {
			newImage = append(newImage, newRow)
		} else {
			blankRow := []string{}
			for range newRow {
				blankRow = append(blankRow, ".")
			}
			newImage = append(newImage, newRow, blankRow)
		}
	}
	return newImage
}

func createGraph(image [][]string) graph.Graph[string, util.Position] {
	g := graph.New(func(v util.Position) string {
		return v.String()
	})
	// add vertexes
	for r, row := range image {
		for c := range row {
			g.AddVertex(util.NewPosition(r, c))
		}
	}
	// add edges
	for r, row := range image {
		for c := range row {
			p := util.NewPosition(r, c)
			for _, dir := range []util.Direction{util.Up, util.Down, util.Left, util.Right} {
				p2 := p.Move(dir, 1)
				g.AddEdge(p.String(), p2.String())
				// if err != nil {
				// 	fmt.Printf("err with p2 %s: %v\n", p2, err)
				// }
			}
		}
	}
	return g
}

func getGalaxyPositions(image [][]string) []util.Position {
	gp := []util.Position{}
	for r, row := range image {
		for c, v := range row {
			if v == "#" {
				p := util.NewPosition(r, c)
				gp = append(gp, p)
			}
		}
	}
	return gp
}

func part1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := stringString(util.ReaderToStrings(f))
	printImage(lines)
	image := expand(lines)
	printImage(image)
	imageGraph := createGraph(image)
	fmt.Println("graph created")
	galaxies := getGalaxyPositions(image)
	fmt.Println("galaxy positions:", galaxies)
	fmt.Println("galaxies:", len(galaxies))
	// now find all shortest paths between them
	shortestPathMap := map[string]int{}
	uniqPairs := []string{}
	for i1, g1 := range galaxies {
		for i2, g2 := range galaxies {
			if i1 != i2 {
				key1 := fmt.Sprintf("%d-%d", i1+1, i2+1)
				key2 := fmt.Sprintf("%d-%d", i2+1, i1+1)
				_, inMap := shortestPathMap[key1]
				if !inMap {
					// find shortest path
					path, err := graph.ShortestPath(imageGraph, g1.String(), g2.String())
					if err != nil {
						panic(err)
					}
					if (i2+1)%100 == 0 {
						fmt.Println(key1, len(path)-1, g1, g2)
					}
					shortestPathMap[key1] = len(path) - 1
					shortestPathMap[key2] = len(path) - 1
					uniqPairs = append(uniqPairs, key1)
				}
			}
		}
	}
	fmt.Println("uniqPairs:", len(uniqPairs))
	lens := []int{}
	for _, k := range uniqPairs {
		lens = append(lens, shortestPathMap[k])
	}
	fmt.Println("sum:", util.Sum(lens...))
}

func main() {
	part1("input.txt")
}
