package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/softwarebygabe/advent/pkg/colors"
	"github.com/softwarebygabe/advent/pkg/util"
	"gonum.org/v1/gonum/stat/combin"
)

func isValid(carts []int, sets []int) bool {
	// fmt.Println(carts, ",", sets)
	// short circuit if 1's don't match
	if util.Sum(carts...) != util.Sum(sets...) {
		return false
	}
	var startIdx int
	for _, s := range sets {
		// fmt.Println(s)
		stack := util.NewStack[int]()
		for i, v := range carts {
			if i < startIdx {
				continue
			}
			// fmt.Println(stack, v)
			if stack.Len() == s {
				// if this current v is not a period, invalid
				if v != 0 {
					return false
				}
				// else we are done with this set
				startIdx = i + 1
				break
			}
			if stack.Len() > 0 && v == 0 {
				// we've started a stack but hit a 0 before the end
				return false
			}
			// if we are at a 1 but stack not full, add to stack
			if v == 1 {
				stack.Push(v)
				startIdx = i + 1
			}
			// if its a 0 ignore
		}
		// fmt.Println("last stack", stack)
		// if we've gone through carts and stack not full, then invalid
		if stack.Len() != s {
			return false
		}
	}
	// if we've gone through set and there are any remaining non-zero carts, invalid
	// fmt.Println(startIdx)
	for i := startIdx; i < len(carts); i++ {
		if carts[i] != 0 {
			return false
		}
	}
	return true
}

func prettySprintCarts(carts []int) string {
	cartStrings := []string{}
	for _, v := range carts {
		s := ""
		switch v {
		case 1:
			s = colors.Sprintf(colors.Red, "%d", v)
		default:
			s = colors.Sprintf(colors.Green, "%d", v)
		}
		cartStrings = append(cartStrings, s)
	}
	return fmt.Sprintf("[%s]", strings.Join(cartStrings, " "))
}

func genAllValidCartesians(slots int, sets []int) [][]int {
	cartLens := make([]int, 0)
	for i := 0; i < slots; i++ {
		cartLens = append(cartLens, 2)
	}
	gen := combin.NewCartesianGenerator(cartLens)
	// Now loop over all products.
	var validCarts [][]int
	var i int
	carts := make([]int, len(cartLens))
	for gen.Next() {
		start := time.Now()
		_ = gen.Product(carts)
		dur := time.Since(start)
		if i%100000 == 0 {
			fmt.Println("gen:", dur, "checking", i, prettySprintCarts(carts))
		}
		if isValid(carts, sets) {
			fmt.Println(i, carts)
			validCarts = append(validCarts, carts)
		}
		i++
	}
	return validCarts
}

func matchesDamage(slots []string, carts []int) bool {
	for i := 0; i < len(slots); i++ {
		s := slots[i]
		c := carts[i]
		switch s {
		case "?":
			continue
		case "#":
			if c != 1 {
				return false
			}
		case ".":
			if c != 0 {
				return false
			}
		default:
			panic("unrecognized slot symbol")
		}
	}
	return true
}

func part1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)

	possibilities := []int{}
	for _, line := range lines {
		fmt.Println("processing line", line)
		line1, line2, _ := strings.Cut(line, " ")
		slots := strings.Split(line1, "")
		numsRaw := strings.Split(line2, ",")
		var nums []int
		for _, n := range numsRaw {
			v, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			nums = append(nums, v)
		}
		validCarts := genAllValidCartesians(len(slots), nums)
		fmt.Println("valid carts", len(validCarts))
		// fmt.Println("valid carts:")
		// for i, v := range validCarts {
		// 	fmt.Println(i+1, v)
		// }
		// fmt.Println("matching carts:")
		matchingCarts := [][]int{}
		for _, v := range validCarts {
			if matchesDamage(slots, v) {
				// fmt.Println(i+1, v)
				matchingCarts = append(matchingCarts, v)
			}
		}
		fmt.Println("matching carts", len(matchingCarts))
		possibilities = append(possibilities, len(matchingCarts))
	}
	fmt.Println(possibilities)
	fmt.Println("sum:", util.Sum(possibilities...))
}

func part2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)

	possibilities := []int{}
	for _, line := range lines {
		fmt.Println("processing line", line)
		line1, line2, _ := strings.Cut(line, " ")
		// unfold lines
		uline1 := line1
		uline2 := line2
		for i := 0; i < 4; i++ {
			uline1 += "?" + line1
			uline2 += "," + line2
		}
		fmt.Println(uline1, uline2)
		slots := strings.Split(uline1, "")
		numsRaw := strings.Split(uline2, ",")

		var nums []int
		for _, n := range numsRaw {
			v, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			nums = append(nums, v)
		}
		validCarts := genAllValidCartesians(len(slots), nums)
		fmt.Println("valid carts", len(validCarts))
		// fmt.Println("valid carts:")
		// for i, v := range validCarts {
		// 	fmt.Println(i+1, v)
		// }
		// fmt.Println("matching carts:")
		matchingCarts := [][]int{}
		for _, v := range validCarts {
			if matchesDamage(slots, v) {
				// fmt.Println(i+1, v)
				matchingCarts = append(matchingCarts, v)
			}
		}
		fmt.Println("matching carts", len(matchingCarts))
		possibilities = append(possibilities, len(matchingCarts))
	}
	fmt.Println(possibilities)
	fmt.Println("sum:", util.Sum(possibilities...))
}

func main() {
	part2("input_test.txt")
}
