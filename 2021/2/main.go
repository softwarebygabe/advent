package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type Instruction struct {
	AimDelta     int
	ForwardDelta int
}

func parseInput(filename string) []Instruction {
	set := []Instruction{}
	util.EvalEachLine(filename, func(line string) {
		splitList := strings.Split(line, " ")
		newInstruction := Instruction{}
		dir := splitList[0]
		delta, err := strconv.Atoi(splitList[1])
		if err != nil {
			panic(err)
		}
		switch dir {
		case "forward":
			newInstruction.ForwardDelta = delta
		case "up":
			newInstruction.AimDelta = delta * -1
		case "down":
			newInstruction.AimDelta = delta
		}
		set = append(set, newInstruction)
	})
	return set
}

type Position struct {
	Depth      int
	Horizontal int
	Aim        int
}

func calculatePosition(instructionSet []Instruction) Position {
	pos := Position{}
	for _, instruction := range instructionSet {
		pos.Aim += instruction.AimDelta
		pos.Horizontal += instruction.ForwardDelta
		pos.Depth += pos.Aim * instruction.ForwardDelta
	}
	return pos
}

func main() {
	subPosition := calculatePosition(parseInput("./input.txt"))
	fmt.Printf("%+v\n", subPosition)
	fmt.Println(subPosition.Depth * subPosition.Horizontal)
}
