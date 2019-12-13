package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseInput(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	orbitStrings := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// parse line here
		fmt.Println("read ->", line)
		orbitStrings = append(orbitStrings, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return orbitStrings
}

func createOrbitTree(orbitStrings []string) *Node {
	root := &Node{id: "COM"}
	// find the next sourceID
	comOrbits := ""
	runThrough := []string{}
	for i, os := range orbitStrings {
		ids := strings.Split(os, ")")
		if ids[0] == "COM" {
			comOrbits = ids[1]
			root.right = &Node{id: comOrbits}
			// rm the COM
			if i == 0 {
				runThrough = orbitStrings[1:]
			} else if i < len(orbitStrings)-1 {
				runThrough = append(orbitStrings[0:i], orbitStrings[i+1:]...)
			} else {
				runThrough = []string{}
			}
		}
	}

	// now do the rest
	for len(runThrough) > 0 {
		runThrough, comOrbits = findAndInsert(runThrough, root, comOrbits)
	}

	return root
}

func findAndInsert(orbitStrings []string, tree *Node, nextSourceID string) ([]string, string) {
	// find the next one we are searching for
	for i, os := range orbitStrings {
		ids := strings.Split(os, ")")
		if ids[0] == nextSourceID {
			insert(tree, ids[0], ids[1])
			if i == 0 {
				orbitStrings = orbitStrings[1:]
			} else if i < len(orbitStrings)-1 {
				orbitStrings = append(orbitStrings[0:i], orbitStrings[i+1:]...)
			} else {
				orbitStrings = []string{}
			}
			return orbitStrings, ids[1]
		}
	}
	// keep it going
	if len(orbitStrings) > 0 {
		return orbitStrings, strings.Split(orbitStrings[0], ")")[0]
	}
	return orbitStrings, ""

}

type Node struct {
	id    string
	right *Node
	left  *Node
}

func find(root *Node, id string) *Node {
	if root == nil {
		return nil
	}
	if root.id == id {
		return root
	}
	findRight := find(root.right, id)
	if findRight != nil {
		return findRight
	}
	findLeft := find(root.left, id)
	if findLeft != nil {
		return findLeft
	}
	return nil
}

func insert(root *Node, id1, id2 string) *Node {
	id1Node := find(root, id1)
	if id1Node == nil {
		panic("we couldn't find the droids you are looking for")
	}

	if id1Node.right == nil {
		id1Node.right = &Node{id: id2}
	} else if id1Node.left == nil {
		id1Node.left = &Node{id: id2}
	} else {
		panic("uhh trying to add to a full binary node")
	}

	return root
}

// countPath counts the number of paths between the root and the Node with destID
func countAllPaths(root *Node) int {
	pathCount := 0
	nodeChan := Walker(root)
	for {
		nodeID, ok := <-nodeChan
		if !ok {
			// no more nodes
			break
		}
		pathCount += getPath(0, root, nodeID)
	}
	return pathCount
}

func getPath(pathCount int, currNode *Node, destID string) int {
	if currNode == nil {
		return 0
	}
	// reach destination return the pathcount
	if currNode.id == destID {
		return pathCount
	}
	lookLeft := getPath(pathCount+1, currNode.left, destID)
	if lookLeft != 0 {
		return lookLeft
	}
	return getPath(pathCount+1, currNode.right, destID)
}

// Walk traverses a tree depth-first,
// sending each Value on a channel.
func Walk(o *Node, ch chan string) {
	if o == nil {
		return
	}
	fmt.Println(o)
	Walk(o.left, ch)
	ch <- o.id
	Walk(o.right, ch)
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func Walker(o *Node) <-chan string {
	ch := make(chan string)
	go func() {
		Walk(o, ch)
		close(ch)
	}()
	return ch
}

func getNodeCount(root *Node) int {
	sum := 0
	c := Walker(root)
	for {
		_, ok := <-c
		if !ok {
			break
		}
		sum++
	}
	return sum
}

func main() {
	fmt.Println("Hello World")

	orbitStrings := parseInput("../test_input1.txt")

	tree := createOrbitTree(orbitStrings)

	// com := &Node{id: "COM"}
	// b := &Node{id: "B"}
	// g := &Node{id: "G"}
	// h := &Node{id: "H"}
	// c := &Node{id: "C"}
	// d := &Node{id: "D"}
	// e := &Node{id: "E"}
	// i := &Node{id: "I"}
	// f := &Node{id: "F"}
	// j := &Node{id: "J"}
	// k := &Node{id: "K"}
	// l := &Node{id: "L"}

	// k.right = l
	// j.right = k
	// e.right = j
	// e.left = f
	// d.right = e
	// d.left = i
	// c.right = d
	// g.right = h
	// b.right = g
	// b.left = c
	// com.right = b

	// tree := com

	fmt.Println("sum ->", countAllPaths(tree))

}
