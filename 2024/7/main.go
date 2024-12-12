package main

import (
	"fmt"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
	"gonum.org/v1/gonum/stat/combin"
)

type Equation struct {
	Result  int
	Numbers []int
}

type operation func(int, int) int

func (e Equation) process(ops []operation) int {
	if len(e.Numbers) < 2 {
		panic("need more than two numbers")
	}
	opsQueue := util.NewQueue[operation]()
	for _, op := range ops {
		opsQueue.Enqueue(op)
	}
	var result int
	var i int
	for i < len(e.Numbers) {
		op, _ := opsQueue.Dequeue()
		if i == 0 {
			result = op(e.Numbers[i], e.Numbers[i+1])
			i++
			i++
		} else {
			result = op(result, e.Numbers[i])
			i++
		}
	}
	return result
}

func add(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

func combine(a, b int) int {
	return util.MustParseInt(fmt.Sprintf("%d%d", a, b))
}

func (e Equation) isValid(operationOptions []operation) bool {
	cartLens := []int{}
	for i := 0; i < len(e.Numbers)-1; i++ {
		cartLens = append(cartLens, len(operationOptions))
	}
	combs := combin.Cartesian(cartLens)
	for _, comb := range combs {
		ops := []operation{}
		for _, combI := range comb {
			ops = append(ops, operationOptions[combI])
		}
		res := e.process(ops)
		if res == e.Result {
			return true
		}
	}
	return false
}

func Part1(filename string) {
	lines, err := util.ReadInput(filename, util.ReaderToStrings)
	if err != nil {
		panic(err)
	}
	equations := []Equation{}
	for _, line := range lines {
		result, numsString, _ := strings.Cut(line, ":")
		nums := strings.Split(strings.Trim(numsString, " "), " ")
		equations = append(equations, Equation{
			Result:  util.MustParseInt(result),
			Numbers: util.StringsToInts(nums),
		})
	}
	var sum int
	for _, eq := range equations {
		valid := eq.isValid([]operation{add, mult})
		fmt.Println(valid, eq)
		if valid {
			sum += eq.Result
		}
	}
	fmt.Println("result:", sum)
}

func Part2(filename string) {
	lines, err := util.ReadInput(filename, util.ReaderToStrings)
	if err != nil {
		panic(err)
	}
	equations := []Equation{}
	for _, line := range lines {
		result, numsString, _ := strings.Cut(line, ":")
		nums := strings.Split(strings.Trim(numsString, " "), " ")
		equations = append(equations, Equation{
			Result:  util.MustParseInt(result),
			Numbers: util.StringsToInts(nums),
		})
	}
	var sum int
	for _, eq := range equations {
		valid := eq.isValid([]operation{add, mult, combine})
		fmt.Println(valid, eq)
		if valid {
			sum += eq.Result
		}
	}
	fmt.Println("result:", sum)
}

func main() {
	Part1("input_ex.txt")
	// Part1("input_1.txt")
	Part2("input_ex.txt")
	Part2("input_1.txt")
}
