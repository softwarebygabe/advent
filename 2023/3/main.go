package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/softwarebygabe/advent/pkg/util"
)

type char struct {
	raw      string
	rawRune  rune
	value    int
	isNumber bool
	isSymbol bool
}

func part1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)
	grid := [][]char{}
	uniq := map[string]bool{}
	for _, line := range lines {
		newLine := []char{}
		for _, r := range line {
			str := string([]rune{r})
			c := char{
				raw:      str,
				rawRune:  r,
				isNumber: unicode.IsNumber(r),
				isSymbol: unicode.IsSymbol(r) || (unicode.IsPunct(r) && str != "."),
			}
			uniq[c.raw] = c.isSymbol
			// fmt.Println(str, c.isSymbol)
			if unicode.IsNumber(r) {
				num, err := strconv.Atoi(str)
				if err != nil {
					panic(err)
				}
				c.value = num
			}
			newLine = append(newLine, c)
		}
		grid = append(grid, newLine)
	}
	for k, v := range uniq {
		fmt.Println(k, v)
	}
	// for _, g := range grid {
	// 	for _, gx := range g {
	// 		fmt.Printf("{%s %v}", gx.raw, gx.isSymbol)
	// 	}
	// 	fmt.Println()
	// }
	// go through grid to find adjacencies
	numbers := [][]char{}
	accounted := map[string]struct{}{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x].isNumber {
				// check if num is adjacent to a symbol
				var isAdjacent bool
				up := y - 1
				down := y + 1
				left := x - 1
				right := x + 1
				// look up
				if up >= 0 {
					// up-up
					if grid[up][x].isSymbol {
						isAdjacent = true
					}
					// up-left
					if left > -1 && grid[up][left].isSymbol {
						isAdjacent = true
					}
					// up-right
					if right < len(grid[y]) && grid[up][right].isSymbol {
						isAdjacent = true
					}
				}
				// look down
				if down < len(grid) {
					// down-down
					if grid[down][x].isSymbol {
						isAdjacent = true
					}
					// down-left
					if left > -1 && grid[down][left].isSymbol {
						isAdjacent = true
					}
					// down-right
					if right < len(grid[y]) && grid[down][right].isSymbol {
						isAdjacent = true
					}
				}
				// look right
				if right < len(grid[y]) {
					if grid[y][right].isSymbol {
						isAdjacent = true
					}
				}
				// look left
				if left >= 0 {
					if grid[y][left].isSymbol {
						isAdjacent = true
					}
				}
				// if is adjacent, collate number on row
				// fmt.Println(grid[y][x].raw, isAdjacent)
				if isAdjacent {
					coord := func(y, x int) string {
						return fmt.Sprintf("(y:%d,x:%d)", y, x)
					}
					if _, ok := accounted[coord(y, x)]; !ok {
						newNumber := []char{}
						xi := x
						// move left until not a number
						for xi > 0 && grid[y][xi].isNumber {
							if grid[y][xi-1].isNumber {
								xi--
							} else {
								break
							}
						}
						// now move right until not a number, collecting all numbers along the way
						for xi < len(grid[y]) && grid[y][xi].isNumber {
							newNumber = append(newNumber, grid[y][xi])
							accounted[coord(y, xi)] = struct{}{}
							xi++
						}
						numbers = append(numbers, newNumber)
					}
				}
			}
		}
	}
	// fmt.Println(numbers)
	allNums := []int{}
	for _, number := range numbers {
		numString := ""
		for _, digit := range number {
			numString += digit.raw
		}
		num, err := strconv.Atoi(numString)
		if err != nil {
			panic(err)
		}
		allNums = append(allNums, num)
	}
	fmt.Println("all nums", allNums)
	fmt.Println("sum:", util.Sum(allNums...))
}

func part2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)
	grid := [][]char{}
	uniq := map[string]bool{}
	for _, line := range lines {
		newLine := []char{}
		for _, r := range line {
			str := string([]rune{r})
			c := char{
				raw:      str,
				rawRune:  r,
				isNumber: unicode.IsNumber(r),
				isSymbol: unicode.IsSymbol(r) || (unicode.IsPunct(r) && str != "."),
			}
			uniq[c.raw] = c.isSymbol
			// fmt.Println(str, c.isSymbol)
			if unicode.IsNumber(r) {
				num, err := strconv.Atoi(str)
				if err != nil {
					panic(err)
				}
				c.value = num
			}
			newLine = append(newLine, c)
		}
		grid = append(grid, newLine)
	}
	for k, v := range uniq {
		fmt.Println(k, v)
	}
	// for _, g := range grid {
	// 	for _, gx := range g {
	// 		fmt.Printf("{%s %v}", gx.raw, gx.isSymbol)
	// 	}
	// 	fmt.Println()
	// }
	// go through grid to find adjacencies
	accounted := map[string]struct{}{}
	gearAdjacent := map[string][][]char{}
	coord := func(y, x int) string {
		return fmt.Sprintf("(y:%d,x:%d)", y, x)
	}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x].isNumber {
				// check if num is adjacent to a symbol
				var isAdjacent bool
				var adjacentChar char
				var adjacentCharCoord string
				up := y - 1
				down := y + 1
				left := x - 1
				right := x + 1
				// look up
				if up >= 0 {
					// up-up
					if grid[up][x].isSymbol {
						isAdjacent = true
						adjacentChar = grid[up][x]
						adjacentCharCoord = coord(up, x)
					}
					// up-left
					if left > -1 && grid[up][left].isSymbol {
						isAdjacent = true
						adjacentChar = grid[up][left]
						adjacentCharCoord = coord(up, left)
					}
					// up-right
					if right < len(grid[y]) && grid[up][right].isSymbol {
						isAdjacent = true
						adjacentChar = grid[up][right]
						adjacentCharCoord = coord(up, right)
					}
				}
				// look down
				if down < len(grid) {
					// down-down
					if grid[down][x].isSymbol {
						isAdjacent = true
						adjacentChar = grid[down][x]
						adjacentCharCoord = coord(down, x)
					}
					// down-left
					if left > -1 && grid[down][left].isSymbol {
						isAdjacent = true
						adjacentChar = grid[down][left]
						adjacentCharCoord = coord(down, left)
					}
					// down-right
					if right < len(grid[y]) && grid[down][right].isSymbol {
						isAdjacent = true
						adjacentChar = grid[down][right]
						adjacentCharCoord = coord(down, right)
					}
				}
				// look right
				if right < len(grid[y]) {
					if grid[y][right].isSymbol {
						isAdjacent = true
						adjacentChar = grid[y][right]
						adjacentCharCoord = coord(y, right)
					}
				}
				// look left
				if left >= 0 {
					if grid[y][left].isSymbol {
						isAdjacent = true
						adjacentChar = grid[y][left]
						adjacentCharCoord = coord(y, left)
					}
				}
				// if is adjacent, collate number on row
				// fmt.Println(grid[y][x].raw, isAdjacent)
				if isAdjacent {
					if _, ok := accounted[coord(y, x)]; !ok {
						newNumber := []char{}
						xi := x
						// move left until not a number
						for xi > 0 && grid[y][xi].isNumber {
							if grid[y][xi-1].isNumber {
								xi--
							} else {
								break
							}
						}
						// now move right until not a number, collecting all numbers along the way
						for xi < len(grid[y]) && grid[y][xi].isNumber {
							newNumber = append(newNumber, grid[y][xi])
							accounted[coord(y, xi)] = struct{}{}
							xi++
						}
						if adjacentChar.raw == "*" {
							fmt.Println("found a gear-adjacent number")
							gearAdjacentNumbers, ok := gearAdjacent[adjacentCharCoord]
							if !ok {
								gearAdjacent[adjacentCharCoord] = [][]char{newNumber}
							} else {
								gearAdjacent[adjacentCharCoord] = append(gearAdjacentNumbers, newNumber)
							}
						}
					}
				}
			}
		}
	}
	digitsToNumber := func(c []char) int {
		numString := ""
		for _, digit := range c {
			numString += digit.raw
		}
		num, err := strconv.Atoi(numString)
		if err != nil {
			panic(err)
		}
		return num
	}
	gearRatios := []int{}
	for k, v := range gearAdjacent {
		if len(v) == 2 {
			fmt.Println("gear qualifies at:", k)
			gearRatios = append(gearRatios, digitsToNumber(v[0])*digitsToNumber(v[1]))
		}
	}
	fmt.Println("gearRatios", gearRatios)
	fmt.Println("sum:", util.Sum(gearRatios...))
}

func main() {
	part2("./input.txt")
}
