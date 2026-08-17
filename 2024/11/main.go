package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/softwarebygabe/advent/pkg/util"
	"golang.org/x/sync/errgroup"
)

type Node struct {
	value int
	right *Node
	left  *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("value=%d leaf=%v", n.value, n.Leaf())
}

func (n *Node) Leaf() bool {
	if n == nil {
		return false
	}
	return n.right == nil && n.left == nil
}

func (n *Node) Blink() {
	// fmt.Println("blink", n)
	if n.value == 0 {
		n.value = 1
		return
	}
	vs := fmt.Sprintf("%d", n.value)
	if len(vs)%2 == 0 {
		// even number of digits
		// split stone
		n.left = &Node{
			value: util.MustParseInt(vs[:len(vs)/2]),
		}
		n.right = &Node{
			value: util.MustParseInt(vs[len(vs)/2:]),
		}
		return
	}
	n.value *= 2024
}

func TraverseLeavesInOrder(n *Node, f func(n *Node)) {
	if n.Leaf() {
		f(n)
		return
	}
	// visit left
	if n.left != nil {
		TraverseLeavesInOrder(n.left, f)
	}
	// visit right
	if n.right != nil {
		TraverseLeavesInOrder(n.right, f)
	}
}

func parseNodes(filename string) []*Node {
	lines, err := util.ReadInput(filename, util.ReaderToStrings)
	if err != nil {
		panic(err)
	}
	nodes := []*Node{}
	for _, char := range strings.Split(lines[0], " ") {
		nodes = append(nodes, &Node{
			value: util.MustParseInt(char),
		})
	}
	return nodes
}

func RunBlink(nodes []*Node) {
	for _, n := range nodes {
		TraverseLeavesInOrder(n, func(n *Node) {
			n.Blink()
		})
	}
}

func RunBlinkParallel(nodes []*Node) {
	var g errgroup.Group
	g.SetLimit(10000)
	for _, n := range nodes {
		n := n
		g.Go(func() error {
			TraverseLeavesInOrder(n, func(n *Node) {
				n.Blink()
			})
			return nil
		})
	}
	g.Wait()
}

func LeafValues(nodes []*Node) []int {
	leaves := []int{}
	for _, n := range nodes {
		TraverseLeavesInOrder(n, func(n *Node) {
			leaves = append(leaves, n.value)
		})
	}
	return leaves
}

func Run(filename string) {
	nodes := parseNodes(filename)
	// numBlinks := 25 // Part 1
	numBlinks := 75 // Part 2
	for i := 0; i < numBlinks; i++ {
		start := time.Now()
		RunBlinkParallel(nodes)
		// fmt.Println(len(LeafValues(nodes)))
		fmt.Printf("RunBlinkParallel(%d) dur: %v\n", i, time.Since(start))
	}
	leaves := LeafValues(nodes)
	fmt.Println(leaves)
	fmt.Println(len(leaves))
}

func main() {
	// Run("input_ex.txt")
	Run("input_1.txt")
}
