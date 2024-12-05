package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type RuleSet map[int][]int

func (r RuleSet) InOrder(n1, n2 int) bool {
	mustBeAfter, ok := r[n2]
	if !ok {
		return true
	}
	for _, v := range mustBeAfter {
		if v == n1 {
			return false
		}
	}
	// valid
	return true
}

type PrinterUpdate []int

func (p PrinterUpdate) Ordered(ruleSet RuleSet) bool {
	for i, n1 := range p {
		for j := i + 1; j < len(p); j++ {
			if !ruleSet.InOrder(n1, p[j]) {
				return false
			}
		}
	}
	return true
}

func (p PrinterUpdate) Middle() int {
	return p[len(p)/2]
}

func (p PrinterUpdate) Sort(ruleSet RuleSet) {
	sort.SliceStable(p, func(i, j int) bool {
		return ruleSet.InOrder(p[i], p[j])
	})
}

func parseLines(lines []string) (RuleSet, []PrinterUpdate) {
	orderRuleMap := map[int][]int{}
	printerUpdates := []PrinterUpdate{}

	var modeTwo bool
	for _, line := range lines {
		if line == "" {
			modeTwo = true
			continue
		}
		if !modeTwo {
			n1String, n2String, _ := strings.Cut(line, "|")
			n1, n2 := util.MustParseInt(n1String), util.MustParseInt(n2String)
			v, ok := orderRuleMap[n1]
			if !ok {
				orderRuleMap[n1] = []int{n2}
			} else {
				orderRuleMap[n1] = append(v, n2)
			}
		} else {
			// modeTwo
			pages := []int{}
			nStrings := strings.Split(line, ",")
			for _, nString := range nStrings {
				pages = append(pages, util.MustParseInt(nString))
			}
			printerUpdates = append(printerUpdates, pages)
		}
	}
	return orderRuleMap, printerUpdates
}

func Part1(fname string) {
	lines, err := util.ReadInput(fname, util.ReaderToStrings)
	if err != nil {
		panic(err)
	}

	ruleSet, printerUpdates := parseLines(lines)
	var sum int
	for _, update := range printerUpdates {
		ordered := update.Ordered(ruleSet)
		fmt.Println(update, ordered)
		if ordered {
			sum += update.Middle()
		}
	}
	fmt.Println("result:", sum)
}

func Part2(fname string) {
	lines, err := util.ReadInput(fname, util.ReaderToStrings)
	if err != nil {
		panic(err)
	}

	ruleSet, printerUpdates := parseLines(lines)
	var sum int
	for _, update := range printerUpdates {
		ordered := update.Ordered(ruleSet)
		if !ordered {
			// sort, take middle, add to sum
			fmt.Println("unsorted", update)
			update.Sort(ruleSet)
			fmt.Println("sorted", update)
			sum += update.Middle()
		}
	}
	fmt.Println("result:", sum)
}

func main() {
	// Part1("input_ex.txt")
	// Part1("input_1.txt")
	// Part2("input_ex.txt")
	Part2("input_1.txt")
}
