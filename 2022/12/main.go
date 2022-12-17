package main

import (
	"fmt"

	"github.com/softwarebygabe/advent/pkg/util"
	"github.com/xlab/treeprint"
)

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
	parent      *node
	children    []*node
	x, y, depth int
	value       rune
	isEnd       bool
	visited     bool
}

func newNode(value rune, x, y int) *node {
	return &node{value: value, children: make([]*node, 0), x: x, y: y}
}

func (n *node) String() string {
	nodeString := fmt.Sprintf("node: x=%d y=%d v=%c", n.x, n.y, n.value)
	// if n.parent != nil {
	// 	nodeString += fmt.Sprintf("\nparent: x=%d y=%d v=%c", n.parent.x, n.parent.y, n.parent.value)
	// } else {
	// 	nodeString += "\nparent: <nil>"
	// }
	return nodeString
}

func (n *node) visit() {
	n.visited = true
}

func (n *node) addChild(cn *node) {
	cn.depth = n.depth + 1
	cn.parent = n
	n.children = append(n.children, cn)
}

func (n *node) eq(n2 *node) bool {
	return n.x == n2.x && n.y == n2.y
}

func parseInput(filepath string) [][]*node {
	result := [][]*node{}
	var y int
	util.EvalEachLine(filepath, func(line string) {
		runes := []*node{}
		for x, r := range line {
			n := newNode(r, x, y)
			if r == 'E' {
				n.isEnd = true
			}
			runes = append(runes, n)
		}
		result = append(result, runes)
		y++
	})
	return result
}

type nodeQueue struct {
	nodes []*node
}

func newNodeQueue() *nodeQueue {
	return &nodeQueue{nodes: make([]*node, 0)}
}

func (nq *nodeQueue) enqueue(n *node) {
	nq.nodes = append(nq.nodes, n)
}

func (nq *nodeQueue) dequeue() *node {
	if len(nq.nodes) < 1 {
		return nil
	}
	n := nq.nodes[0]
	nq.nodes = nq.nodes[1:]
	return n
}

func (nq *nodeQueue) empty() bool {
	return len(nq.nodes) == 0
}

func processTo(board [][]*node, pn *node, x, y int) *node {
	to := board[y][x]
	if pn.parent != nil && pn.parent.eq(to) {
		return nil
	}
	if to.visited {
		return nil
	}
	if canMoveTo(pn.value, to.value) {
		pn.addChild(to)
		return to
	}
	return nil
}

func printNode(t treeprint.Tree, n *node) {
	b := t.AddBranch(n)
	for _, cn := range n.children {
		printNode(b, cn)
	}
}

func PrintTree(root *node) {
	tree := treeprint.New()
	tree.SetValue(root.String())
	for _, cn := range root.children {
		printNode(tree, cn)
	}
	fmt.Println(tree.String())
}

func part1() {
	board := parseInput("input_test.txt")
	root := board[0][0]
	// root.addChild(newNode(board[0][1]))
	// root.addChild(newNode(board[1][0]))
	queue := newNodeQueue()
	queue.enqueue(root)
	var answer int
	for !queue.empty() {
		PrintTree(root)
		// currCycle++
		currParent := queue.dequeue()
		currParent.visit()
		// look up, down, left, right
		upY := currParent.y - 1
		downY := currParent.y + 1
		if -1 < upY {
			cn := processTo(board, currParent, currParent.x, upY)
			if cn != nil && cn.isEnd {
				answer = cn.depth
				break
			}
		}
		if downY < len(board) {
			cn := processTo(board, currParent, currParent.x, downY)
			if cn != nil && cn.isEnd {
				answer = cn.depth
				break
			}
		}
		leftX := currParent.x - 1
		rightX := currParent.x + 1
		if -1 < leftX {
			cn := processTo(board, currParent, leftX, currParent.y)
			if cn != nil && cn.isEnd {
				answer = cn.depth
				break
			}
		}
		if rightX < len(board[currParent.y]) {
			cn := processTo(board, currParent, rightX, currParent.y)
			if cn != nil && cn.isEnd {
				answer = cn.depth
				break
			}
		}
		// fmt.Println(currParent)
		for _, cn := range currParent.children {
			queue.enqueue(cn)
		}
	}
	fmt.Println("steps:", answer)
}

func main() {
	part1()
}
