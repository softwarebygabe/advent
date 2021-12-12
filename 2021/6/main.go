package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

func parseInput(filename string) []int {
	fishPool := []int{}
	util.EvalEachLine(filename, func(line string) {
		nums := strings.Split(line, ",")
		for _, num := range nums {
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			fishPool = append(fishPool, i)
		}
	})
	return fishPool
}

func main() {
	fishPool := parseInput("./input.txt")
	// init map
	amtMap := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}
	for _, fish := range fishPool {
		sum, ok := amtMap[fish]
		if !ok {
			amtMap[fish] = 1
		} else {
			amtMap[fish] = sum + 1
		}
	}
	fmt.Println(amtMap)
	// increment days
	for day := 1; day <= 256; day++ {
		newMap := map[int]int{
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
			5: 0,
			6: 0,
			7: 0,
			8: 0,
		}
		for k := range amtMap {
			if k != 0 && k != 7 {
				newMap[k-1] = amtMap[k]
			} else {
				// 6's are 0's + 7's
				newMap[6] = amtMap[0] + amtMap[7]
				// 8's are 0's
				newMap[8] = amtMap[0]
			}
		}
		amtMap = newMap
		fmt.Println(amtMap)
	}
	var total int
	for _, v := range amtMap {
		total += v
	}
	// how many fish are there now?
	fmt.Println("number of fish in pool:", total)
}
