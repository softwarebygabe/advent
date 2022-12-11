package main

import (
	"fmt"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type clock struct {
	hook       func(cycle int)
	cycleCount int
}

func newClock(hook func(cycle int)) *clock {
	return &clock{hook: hook, cycleCount: 1}
}

func (c *clock) cycle() {
	c.hook(c.cycleCount)
	c.cycleCount++
}

type cpu struct {
	register int
	clock    *clock
}

func newCPU() *cpu {
	return &cpu{register: 1}
}

func (c *cpu) setClock(cl *clock) {
	c.clock = cl
}

func (c *cpu) noop() {
	c.clock.cycle()
}

func (c *cpu) addx(n int) {
	c.clock.cycle()
	c.clock.cycle()
	c.register += n
}

func (c *cpu) runProgram(instructions []string) {
	for _, instruction := range instructions {
		elems := strings.Split(instruction, " ")
		switch elems[0] {
		case "noop":
			c.noop()
		case "addx":
			c.addx(util.MustParseInt(elems[1]))
		}
	}
}

type crt struct {
	rowLength int
	rows      []string
	currRow   string
	currPos   int
}

func newCRT() *crt {
	return &crt{
		rowLength: 40,
		rows:      make([]string, 0),
	}
}

func (c *crt) flush() {
	c.rows = append(c.rows, c.currRow)
	c.currRow = ""
	c.currPos = 0
}

func (c *crt) drawChar(spritePos int) {
	if c.currPos < c.rowLength {
		spritePositions := []int{spritePos - 1, spritePos, spritePos + 1}
		for _, sp := range spritePositions {
			if c.currPos == sp {
				c.currRow += "#"
				c.currPos++
				return
			}
		}
		c.currRow += " "
		c.currPos++
	} else {
		c.rows = append(c.rows, c.currRow)
		c.currRow = ""
		c.currPos = 0
		c.drawChar(spritePos)
	}
}

func (c *crt) String() string {
	return strings.Join(c.rows, "\n")
}

func parseInput(filepath string) []string {
	results := []string{}
	util.EvalEachLine(filepath, func(line string) {
		results = append(results, line)
	})
	return results
}

func part1() {
	cpu := newCPU()
	signals := []int{}
	currstep := 0
	clock := newClock(func(cycle int) {
		fmt.Println("during cycle:", cycle, "cpu.register:", cpu.register)
		if cycle == 20 || cycle-currstep == 40 {
			signals = append(signals, cpu.register*cycle)
			currstep = cycle
		}
	})
	cpu.setClock(clock)
	instructions := parseInput("input.txt")
	cpu.runProgram(instructions)
	fmt.Println(signals)
	var sum int
	for _, v := range signals {
		sum += v
	}
	fmt.Println("result:", sum)
}

func part2() {
	cpu := newCPU()
	crt := newCRT()
	clock := newClock(func(cycle int) {
		crt.drawChar(cpu.register)
	})
	cpu.setClock(clock)
	instructions := parseInput("input.txt")
	cpu.runProgram(instructions)
	crt.flush()
	fmt.Println(crt)
}

func main() {
	part2()
}
