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
	// parentID: [childID, childID]
	instructionMap := map[string][]string{}
	for _, orbitString := range orbitStrings {
		osSplit := strings.Split(orbitString, ")")
		parentID := osSplit[0]
		childID := osSplit[1]
		childIDs, inMap := instructionMap[parentID]
		if !inMap {
			instructionMap[parentID] = []string{childID}
		} else {
			instructionMap[parentID] = append(childIDs, childID)
		}
	}

	// now start with the COM as the root and walk map to create tree
	root := &Node{id: "COM"}
	parentQueue := append([]string{"COM"})
	for i := 0; i < len(parentQueue); i++ {
		currentParentID := parentQueue[i]
		parentNode := find(root, currentParentID)
		childIDs := instructionMap[currentParentID]
		for _, childID := range childIDs {
			if parentNode.right != nil && parentNode.left != nil {
				panic("this node is full oh no!!")
			}
			if parentNode.right == nil {
				parentNode.right = &Node{id: childID}
			} else if parentNode.left == nil {
				parentNode.left = &Node{id: childID}
			}
			parentQueue = append(parentQueue, childID)
		}
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

// Node ...
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

func getPath(path []string, currNode *Node, destID string) []string {
	if currNode == nil {
		return []string{}
	}
	// reach destination return the pathcount
	if currNode.id == destID {
		return path
	}
	lookLeft := getPath(append(path, currNode.id), currNode.left, destID)
	if len(lookLeft) > 0 {
		return lookLeft
	}
	return getPath(append(path, currNode.id), currNode.right, destID)
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

	orbitStrings := parseInput("../input.txt")

	tree := createOrbitTree(orbitStrings)

	pathToSanta := getPath([]string{}, tree, "SAN")
	pathToYou := getPath([]string{}, tree, "YOU")

	fmt.Println("path ->>", pathToSanta)
	fmt.Println("path ->>", pathToYou)

	// get intersection nodeID
	intersectionID := ""
	sI := 0
	yI := 0
	for i, sID := range pathToSanta {
		for j, yID := range pathToYou {
			if sID == yID {
				intersectionID = sID
				sI = i
				yI = j
			}
		}
	}

	santaPathCount := len(pathToSanta[sI+1:])
	youPathCount := len(pathToYou[yI+1:])

	fmt.Println("intersection point ->>", intersectionID)
	fmt.Println("orbit transfer amt ->>", santaPathCount+youPathCount)
}
