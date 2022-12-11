package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/softwarebygabe/advent/pkg/util"
)

type monkey struct {
	items        []int
	argA         string
	argB         string
	operation    func(a, b int) int
	divisibleBy  int
	ifTrue       *monkey
	ifFalse      *monkey
	inspectCount int
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func (m *monkey) getOpVal(arg string, old int) int {
	if arg == "old" {
		return old
	}
	return util.MustParseInt(arg)
}

func correction(v int) int {
	v2 := float64(v) / 3.0
	return int(math.Floor(v2))
}

func (m *monkey) playTurn() {
	for _, item := range m.items {
		// op
		a := m.getOpVal(m.argA, item)
		b := m.getOpVal(m.argB, item)
		new := m.operation(a, b)
		// correction
		new = correction(new)
		// test
		if new%m.divisibleBy == 0 {
			m.ifTrue.catch(new)
		} else {
			m.ifFalse.catch(new)
		}
	}
}

func (m *monkey) catch(item int) {
	m.items = append(m.items, item)
}

// ByInspectCount implements sort.Interface for []monkey based on
// the inspectCount field.
type ByInspectCount []*monkey

func (a ByInspectCount) Len() int           { return len(a) }
func (a ByInspectCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByInspectCount) Less(i, j int) bool { return a[i].inspectCount < a[j].inspectCount }

func parseInput(filepath string) []*monkey {
	return nil
}

func part1() {
	monkeys := parseInput("input_test.txt")
	for i := 0; i < 20; i++ {
		// round
		for _, monkey := range monkeys {
			monkey.playTurn()
		}
	}
	sort.Sort(ByInspectCount(monkeys))
	a := monkeys[len(monkeys)-1].inspectCount
	b := monkeys[len(monkeys)-2].inspectCount
	fmt.Println("result:", multiply(a, b))
}

func main() {
	part1()
}
