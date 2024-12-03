package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

func newValidCharStack() *util.Stack[string] {
	v := util.NewStack[string]()
	v.Push(")")
	v.Push(",")
	v.Push("(")
	return v
}

func process(chars []string) int {
	vcharStack := newValidCharStack()
	mulStack := util.NewStack[string]()
	var n1String string
	var n2String string
	muls := []int{}
	for _, char := range chars {
		fmt.Println("char:", char)
		fmt.Println("mulStack Len:", mulStack.Len())
		fmt.Println("vcharStack Len:", vcharStack.Len())
		fmt.Println("n1String:", n1String)
		fmt.Println("n2String:", n2String)
		switch mulStack.Len() {
		case 0:
			// looking for "m"
			if char == "m" {
				mulStack.Push(char)
			}
			continue
		case 1:
			// looking for "u"
			if char == "u" {
				mulStack.Push(char)
				continue
			}
			mulStack.Reset() // reset if not "u"
		case 2:
			// looking for "l"
			if char == "l" {
				mulStack.Push(char)
				continue
			}
			mulStack.Reset() // reset if not "l"
		case 3:
			// use vcharStack, and parse nums
			resetInner := func() {
				mulStack.Reset()
				vcharStack = newValidCharStack()
				n1String = ""
				n2String = ""
			}
			switch vcharStack.Len() {
			case 3:
				// check if we are at "("
				if char == "(" {
					vcharStack.Pop()
					continue
				}
				// reset
				resetInner()
				continue
			case 2:
				// check if we are at the vchar ","
				if char == "," {
					// move on
					vcharStack.Pop()
					continue
				}
				// check if we have a number
				_, err := strconv.Atoi(char)
				if err != nil {
					// not a number, reset
					resetInner()
					continue
				}
				// is a number
				n1String += char
				continue
			case 1:
				// check if we are at ")"
				if char == ")" {
					// perform math
					n1, err := strconv.Atoi(n1String)
					if err != nil {
						panic(err)
					}
					n2, err := strconv.Atoi(n2String)
					if err != nil {
						panic(err)
					}
					muls = append(muls, n1*n2)
					// reset
					resetInner()
					continue
				}
				// check if we have a number
				_, err := strconv.Atoi(char)
				if err != nil {
					// not a number, reset
					resetInner()
					continue
				}
				// is a number
				n2String += char
				continue
			default:
				panic("unreachable")
			}
		default:
			panic("unreachable")
		}
	}

	return util.Sum(muls...)
}

func Part1(fname string) {
	lines, err := util.ReadInput[[]string](fname, util.ReaderToStrings)
	if err != nil {
		panic(err)
	}
	var fullLine string
	for _, line := range lines {
		fullLine += line
	}

	allChars := strings.Split(fullLine, "")
	fmt.Println("result:", process(allChars))
}

func Part2(fname string) {
	lines, err := util.ReadInput[[]string](fname, util.ReaderToStrings)
	if err != nil {
		panic(err)
	}
	var fullLine string
	for _, line := range lines {
		fullLine += line
	}
	splitByDont := strings.Split(fullLine, "don't()")
	sums := []int{process(strings.Split(splitByDont[0], ""))}
	for i := 1; i < len(splitByDont); i++ {
		splitByDo := strings.Split(splitByDont[i], "do()")
		for idx, section := range splitByDo {
			if idx > 0 {
				sums = append(sums, process(strings.Split(section, "")))
			}
		}
	}
	fmt.Println("result:", util.Sum(sums...))
}

func main() {
	// Part1("input_ex.txt")
	// Part1("input_1.txt")
	// Part2("input_ex2.txt")
	Part2("input_1.txt")
}
