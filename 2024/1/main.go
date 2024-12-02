package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

func Part1() {
	filename := "input_part1.txt"
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)
	leftList := make([]int, 0, len(lines))
	rightList := make([]int, 0, len(lines))
	for _, line := range lines {
		l, r, _ := strings.Cut(line, "   ")
		li, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		ri, err := strconv.Atoi(r)
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, li)
		rightList = append(rightList, ri)
	}
	// sort both lists
	sort.Ints(leftList)
	sort.Ints(rightList)
	// now add distances
	totDist := 0
	for i := 0; i < len(leftList); i++ {
		totDist += int(math.Abs(float64(leftList[i]) - float64(rightList[i])))
	}
	fmt.Println("total distance:", totDist)
}

func Part2() {
	filename := "input_part1.txt"
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)
	leftList := make([]int, 0, len(lines))
	rightList := make([]int, 0, len(lines))
	for _, line := range lines {
		l, r, _ := strings.Cut(line, "   ")
		li, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		ri, err := strconv.Atoi(r)
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, li)
		rightList = append(rightList, ri)
	}
	seenCount := make(map[int]int)
	for _, l := range leftList {
		_, ok := seenCount[l]
		if !ok {
			count := 0
			for _, r := range rightList {
				if r == l {
					count++
				}
			}
			seenCount[l] = count
		}
	}
	totScore := 0
	for k, v := range seenCount {
		totScore += k * v
	}
	fmt.Println("total score:", totScore)
}

func main() {
	// Part1()
	Part2()
}
