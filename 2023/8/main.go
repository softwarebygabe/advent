package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type node struct {
	name  string
	right *node
	left  *node
}

func (n *node) String() string {
	return n.name
}

func part1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)
	instructions := []string{}
	network := make(map[string]*node)
	ensureNode := func(name string) *node {
		n, ok := network[name]
		if !ok {
			newNode := &node{name: name}
			network[name] = newNode
			return newNode
		}
		return n
	}
	var rootNode *node
	for idx, line := range lines {
		// fmt.Println(network)
		if idx == 0 {
			instructions = strings.Split(line, "")
			continue
		}
		if line != "" {
			nodeName, linkedRaw, ok := strings.Cut(line, " = ")
			if !ok {
				panic("could not cut")
			}
			linked := strings.Split(strings.TrimPrefix(
				strings.TrimSuffix(linkedRaw, ")"),
				"(",
			), ", ")
			n := ensureNode(nodeName)
			if rootNode == nil && nodeName == "AAA" {
				rootNode = n
			}
			nL := ensureNode(linked[0])
			nR := ensureNode(linked[1])
			n.left = nL
			n.right = nR
		}
	}
	var reachedZZZ bool
	var moveCount int
	var cursor int
	currentNode := rootNode
	for !reachedZZZ {
		switch instructions[cursor] {
		case "R":
			currentNode = currentNode.right
		case "L":
			currentNode = currentNode.left
		default:
			panic("cannot interpret instruction")
		}
		moveCount++
		cursor++
		if cursor == len(instructions) {
			cursor = 0
		}
		if currentNode.name == "ZZZ" {
			reachedZZZ = true
		}
	}
	fmt.Println("moves:", moveCount)
}

func part2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)
	instructions := []string{}
	network := make(map[string]*node)
	ensureNode := func(name string) *node {
		n, ok := network[name]
		if !ok {
			newNode := &node{name: name}
			network[name] = newNode
			return newNode
		}
		return n
	}
	for idx, line := range lines {
		// fmt.Println(network)
		if idx == 0 {
			instructions = strings.Split(line, "")
			continue
		}
		if line != "" {
			nodeName, linkedRaw, ok := strings.Cut(line, " = ")
			if !ok {
				panic("could not cut")
			}
			linked := strings.Split(strings.TrimPrefix(
				strings.TrimSuffix(linkedRaw, ")"),
				"(",
			), ", ")
			n := ensureNode(nodeName)
			nL := ensureNode(linked[0])
			nR := ensureNode(linked[1])
			n.left = nL
			n.right = nR
		}
	}
	var reachedAllZs bool
	var moveCount int
	var cursor int
	rootNodes := []*node{}
	for k, v := range network {
		fmt.Println(k, v)
		if k[len(k)-1] == 'A' {
			rootNodes = append(rootNodes, v)
		}
	}
	fmt.Println(rootNodes)
	// get the root nodes
	for !reachedAllZs {
		checkNodes := []*node{}
		instr := instructions[cursor]
		for i := 0; i < len(rootNodes); i++ {
			switch instr {
			case "R":
				checkNodes = append(checkNodes, rootNodes[i].right)
			case "L":
				checkNodes = append(checkNodes, rootNodes[i].left)
			default:
				panic("cannot interpret instruction")
			}
		}
		// fmt.Println("check nodes", checkNodes)
		moveCount++
		cursor++
		if cursor == len(instructions) {
			cursor = 0
		}
		// check nodes
		var notZ bool
		for _, n := range checkNodes {
			if n.name[len(n.name)-1] != 'Z' {
				notZ = true
				break
			}
		}
		if !notZ {
			reachedAllZs = true
		}
		if moveCount%10000000 == 0 {
			fmt.Println(instr)
			fmt.Println(checkNodes)
		}
		rootNodes = checkNodes
	}
	fmt.Println("moves:", moveCount)
}

func main() {
	part2("input.txt")
}
