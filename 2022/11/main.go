package main

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type monkey struct {
	items        []*big.Int
	argA         string
	argB         string
	operation    func(a, b *big.Int) *big.Int
	divisibleBy  *big.Int
	ifTrue       *monkey
	ifFalse      *monkey
	inspectCount int
}

func (m *monkey) String() string {
	return fmt.Sprint(m.items)
}

func add(a, b *big.Int) *big.Int {
	return a.Add(a, b)
}

func multiply(a, b *big.Int) *big.Int {
	return a.Mul(a, b)
}

func (m *monkey) getOpVal(arg string, old *big.Int) *big.Int {
	if arg == "old" {
		return old
	}
	return big.NewInt(int64(util.MustParseInt(arg)))
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
		// new = correction(new)
		// test

		if new.Mod(new, m.divisibleBy).Cmp(big.NewInt(0)) == 0 {
			m.ifTrue.catch(new)
		} else {
			m.ifFalse.catch(new)
		}
		m.inspectCount++
	}
	m.items = []*big.Int{}
}

func (m *monkey) catch(item *big.Int) {
	m.items = append(m.items, item)
}

func printMonkeys(monkeys []*monkey) {
	for idx, monkey := range monkeys {
		fmt.Println("Monkey", idx, "inspected items", monkey.inspectCount, "times.")
	}
}

// ByInspectCount implements sort.Interface for []monkey based on
// the inspectCount field.
type ByInspectCount []*monkey

func (a ByInspectCount) Len() int           { return len(a) }
func (a ByInspectCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByInspectCount) Less(i, j int) bool { return a[i].inspectCount < a[j].inspectCount }

func parseInput(filepath string) []*monkey {
	monkeys := []*monkey{}
	util.EvalEachLine(filepath, func(line string) {
		if strings.Contains(line, "Monkey") {
			monkeys = append(monkeys, &monkey{})
		}
	})
	fmt.Println(len(monkeys))
	var currMonkeyIdx int
	util.EvalEachLine(filepath, func(line string) {
		if strings.Contains(line, "Monkey") {
			return
		}
		if line == "" {
			currMonkeyIdx++
			return
		}
		currMonkey := monkeys[currMonkeyIdx]
		switch strings.TrimSpace(strings.Split(line, ":")[0]) {
		case "Starting items":
			parse := strings.Split(line, ": ")[1]
			elems := strings.Split(parse, ", ")
			for _, elem := range elems {
				currMonkey.catch(big.NewInt(int64(util.MustParseInt(elem))))
			}
		case "Operation":
			parse := strings.Split(line, ": new = ")[1]
			if strings.Contains(parse, "+") {
				elems := strings.Split(parse, " + ")
				currMonkey.argA = elems[0]
				currMonkey.argB = elems[1]
				currMonkey.operation = add
			}
			if strings.Contains(parse, "*") {
				elems := strings.Split(parse, " * ")
				currMonkey.argA = elems[0]
				currMonkey.argB = elems[1]
				currMonkey.operation = multiply
			}
		case "Test":
			parse := strings.Split(line, ": divisible by ")[1]
			currMonkey.divisibleBy = big.NewInt(int64(util.MustParseInt(parse)))
		case "If true":
			parse := strings.Split(line, ": throw to monkey ")[1]
			idx := util.MustParseInt(parse)
			currMonkey.ifTrue = monkeys[idx]
		case "If false":
			parse := strings.Split(line, ": throw to monkey ")[1]
			idx := util.MustParseInt(parse)
			currMonkey.ifFalse = monkeys[idx]

		default:
			panic("unable to parse line: " + line)
		}
	})
	printMonkeys(monkeys)
	return monkeys
}

func part1() {
	monkeys := parseInput("input.txt")
	for i := 0; i < 20; i++ {
		fmt.Println("Round", i+1)
		// round
		for _, monkey := range monkeys {
			monkey.playTurn()
		}
		printMonkeys(monkeys)
	}
	sort.Sort(ByInspectCount(monkeys))
	a := monkeys[len(monkeys)-1].inspectCount
	b := monkeys[len(monkeys)-2].inspectCount
	fmt.Println("result:", a*b)
}

func part2() {
	monkeys := parseInput("input_test.txt")
	for i := 0; i < 10000; i++ {
		// fmt.Println("Round", i+1)
		// round
		for _, monkey := range monkeys {
			monkey.playTurn()
		}
		if i+1 == 1 || i+1 == 20 || (i+1)%1000 == 0 {
			fmt.Println("Round", i+1)
			printMonkeys(monkeys)
		}
	}
	sort.Sort(ByInspectCount(monkeys))
	a := monkeys[len(monkeys)-1].inspectCount
	b := monkeys[len(monkeys)-2].inspectCount
	fmt.Println("result:", a*b)
}

func main() {
	part2()
}
