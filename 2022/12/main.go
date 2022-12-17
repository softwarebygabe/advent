package main

import "github.com/softwarebygabe/advent/pkg/util"

var heights = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p',
	'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

func genMap(letters []rune) map[rune]int {
	result := make(map[rune]int)
	for i, l := range letters {
		result[l] = i
	}
	return result
}

var heightMap = genMap(heights)

func diff(a, b rune) int {
	return heightMap[b] - heightMap[a]
}

func canMoveTo(from, to rune) bool {
	delta := diff(from, to)
	return delta <= 1
}

// construct a tree from the board representing all paths that can be taken from 'S'
// stop when you reach 'E'
// calculate branch length to 'E' (add node depth to nodes, this should be answer)

type node struct {
	parent   *node
	children []*node
	depth    int
	value    rune
}

func newNode(value rune) *node {
	return &node{value: value, children: make([]*node, 0)}
}

func (n *node) addChild(cn *node) {
	cn.depth = n.depth + 1
	cn.parent = n
	n.children = append(n.children, cn)
}

func parseInput(filepath string) [][]rune {
	result := [][]rune{}
	util.EvalEachLine(filepath, func(line string) {
		runes := []rune{}
		for _, r := range line {
			runes = append(runes, r)
		}
		result = append(result, runes)
	})
	return result
}

func part1() {}

func main() {
	part1()
}
