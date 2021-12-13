package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

func median(nums []int) int {
	// sort the list ascending
	sort.Ints(nums)
	fmt.Println(nums)
	if len(nums)%2 != 0 {
		// if not even take the middle
		return nums[len(nums)/2]
	}
	upperF := math.Ceil(float64(len(nums)) / 2)
	upper := int(upperF)
	lower := upper - 1
	return (nums[upper] + nums[lower]) / 2
}

func Part1() {
	input := []int{}
	util.EvalEachLine("./input.txt", func(line string) {
		splitList := strings.Split(line, ",")
		for _, numS := range splitList {
			num, err := strconv.Atoi(numS)
			if err != nil {
				panic(err)
			}
			input = append(input, num)
		}
	})
	median := median(input)
	fmt.Println(median)
	var fuelSum int
	for _, num := range input {
		fuelF := math.Abs(float64(num) - float64(median))
		fuelSum += int(fuelF)
	}
	fmt.Println(fuelSum)
}

func mean(nums []int) float64 {
	var sum float64
	for _, num := range nums {
		sum += float64(num)
	}
	avg := sum / float64(len(nums))
	return avg
}

func main() {
	input := []int{}
	util.EvalEachLine("./input.txt", func(line string) {
		splitList := strings.Split(line, ",")
		for _, numS := range splitList {
			num, err := strconv.Atoi(numS)
			if err != nil {
				panic(err)
			}
			input = append(input, num)
		}
	})
	meanF := mean(input)
	fmt.Println(meanF)
	mean := int(math.Ceil(meanF)) - 1
	fmt.Println(mean)
	var fuelSum int
	for _, num := range input {
		diff := int(math.Abs(float64(num) - float64(mean)))
		for i := 1; i <= diff; i++ {
			fuelSum += i
		}
	}
	fmt.Println(fuelSum)
}
