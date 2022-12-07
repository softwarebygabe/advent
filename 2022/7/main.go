package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
	"github.com/xlab/treeprint"
)

type dir struct {
	name   string
	size   int
	parent *dir
	dirs   []*dir
	files  []*file
}

func newDir(name string) *dir {
	return &dir{name: name}
}

func (d *dir) addDir(d2 *dir) {
	d2.parent = d
	d.dirs = append(d.dirs, d2)
}

func (d *dir) addFile(f *file) {
	d.files = append(d.files, f)
}

type file struct {
	name string
	size int
}

func newFile(name string, size int) *file {
	return &file{name: name, size: size}
}

func (d *dir) Print(t treeprint.Tree) {
	branch := t.AddBranch(fmt.Sprintf("%s (dir, size=%d)", d.name, d.size))
	for _, childFiles := range d.files {
		branch.AddNode(childFiles.String())
	}
	for _, childDir := range d.dirs {
		childDir.Print(branch)
	}
}

func (f *file) String() string {
	return fmt.Sprintf("%s (file, size=%d)", f.name, f.size)
}

func (d *dir) calcSizes() int {
	var dirSize int
	for _, childFile := range d.files {
		dirSize += childFile.size
	}
	for _, childDir := range d.dirs {
		dirSize += childDir.calcSizes()
	}
	d.size = dirSize
	return dirSize
}

func traverseDirs(root *dir, visit func(d *dir)) {
	visit(root)
	for _, d := range root.dirs {
		traverseDirs(d, visit)
	}
}

func getAllDirsBySizeLimit(root *dir, limit int) []*dir {
	underLimit := []*dir{}
	traverseDirs(root, func(d *dir) {
		if d.size <= limit {
			underLimit = append(underLimit, d)
		}
	})
	return underLimit
}

func getAllDirsBySizeMinimum(root *dir, min int) []*dir {
	aboveMin := []*dir{}
	traverseDirs(root, func(d *dir) {
		if min <= d.size {
			aboveMin = append(aboveMin, d)
		}
	})
	return aboveMin
}

func parseInput(filepath string) *dir {
	var root *dir
	var currDir *dir
	currLSLines := []string{}
	processLSLines := func(currDir *dir, lines []string) {
		for _, line := range lines {
			words := strings.Split(line, " ")
			size := words[0]
			name := words[1]
			if size != "dir" {
				// file
				currDir.addFile(newFile(name, util.MustParseInt(size)))
			} else {
				// dir
				currDir.addDir(newDir(name))
			}
		}
	}
	util.EvalEachLine(filepath, func(line string) {
		if line[0] == '$' {
			// cmd
			// process any ls lines ...
			if currDir != nil {
				processLSLines(currDir, currLSLines)
				currLSLines = []string{}
			}
			// clear ls lines
			words := strings.Split(line, " ")
			cmd := words[1]
			switch cmd {
			case "cd":
				dirName := words[2]
				switch dirName {
				case "..":
					currDir = currDir.parent
				default:
					if root == nil {
						root = newDir(dirName)
						currDir = root
					} else {
						for _, child := range currDir.dirs {
							if child.name == dirName {
								currDir = child
							}
						}
					}
				}
			case "ls":
				// nothing
			}
		} else {
			// feed to ls
			currLSLines = append(currLSLines, line)
		}
	})
	processLSLines(currDir, currLSLines)
	currLSLines = []string{}
	return root
}

func part1() {
	root := parseInput("input.txt")
	root.calcSizes()
	tree := treeprint.New()
	root.Print(tree)
	fmt.Println(tree.String())
	underLimit := getAllDirsBySizeLimit(root, 100000)
	fmt.Println(underLimit)
	var sum int
	for _, d := range underLimit {
		sum += d.size
	}
	fmt.Println("result:", sum)
}

func part2() {
	root := parseInput("input.txt")
	root.calcSizes()
	tree := treeprint.New()
	root.Print(tree)
	fmt.Println(tree.String())

	totalSpace := 70000000
	desiredUnused := 30000000

	totalUsedSpace := root.size
	currUnusedSpace := totalSpace - totalUsedSpace

	spaceNeededToBeFreed := desiredUnused - currUnusedSpace

	candidatesForDeletion := getAllDirsBySizeMinimum(root, spaceNeededToBeFreed)
	// sort
	sizes := []int{}
	for _, d := range candidatesForDeletion {
		sizes = append(sizes, d.size)
	}
	sort.Ints(sizes)
	fmt.Println("result:", sizes[0])

}

func main() {
	part2()
}
