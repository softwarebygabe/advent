package main

import (
	"fmt"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

var priorities = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
	"q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
	"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

func getPriority(v string) int {
	for idx, c := range priorities {
		if c == v {
			return idx + 1
		}
	}
	panic("can not find priority")
}

func parseInput(filename string) []string {
	backpacks := []string{}
	util.EvalEachLine(filename, func(line string) {
		backpacks = append(backpacks, line)
	})
	return backpacks
}

func findCommonChar(backpack string) string {
	// fmt.Println(backpack)
	charList := strings.Split(backpack, "")
	c1 := charList[0 : len(charList)/2]
	c2 := charList[len(charList)/2:]
	// fmt.Println(c1)
	// fmt.Println(c2)
	charMap := map[string]struct{}{}
	for _, c := range c1 {
		charMap[c] = struct{}{}
	}
	for _, c := range c2 {
		_, inMap := charMap[c]
		if inMap {
			return c
		}
	}
	panic("can not find common char")
}

func part1() {
	backpacks := parseInput("input.txt")
	chars := []string{}
	for _, b := range backpacks {
		chars = append(chars, findCommonChar(b))
	}
	// fmt.Println(chars)
	var sum int
	for _, c := range chars {
		sum += getPriority(c)
	}
	fmt.Println("sum:", sum)
}

func allTrue(l ...bool) bool {
	for _, b := range l {
		if !b {
			return false
		}
	}
	return true
}

func findBadge(backpacks ...string) string {
	charMap := map[string][]bool{}
	for idx, b := range backpacks {
		bChars := strings.Split(b, "")
		for _, c := range bChars {
			v, inMap := charMap[c]
			if !inMap {
				v = make([]bool, len(backpacks))
				v[idx] = true
				charMap[c] = v
			} else {
				v[idx] = true
				charMap[c] = v
			}
		}
	}
	// fmt.Println(charMap)
	for k, v := range charMap {
		if allTrue(v...) {
			return k
		}
	}
	panic("cannot find badge")
}

func part2() {
	backpacks := parseInput("input.txt")
	badges := []string{}
	for i := 2; i < len(backpacks); i += 3 {
		b1 := backpacks[i-2]
		b2 := backpacks[i-1]
		b3 := backpacks[i]
		badges = append(badges, findBadge(b1, b2, b3))
	}
	var sum int
	for _, c := range badges {
		sum += getPriority(c)
	}
	fmt.Println("sum:", sum)
}

func main() {
	part2()
}
