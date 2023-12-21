package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

func Part1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	strs := util.ReaderToStrings(f)
	fmt.Println(strs)
	// filter to just digits
	digits := [][]int{}
	for _, line := range strs {
		fmt.Println(line)
		lStrs := strings.Split(line, "")
		lDigits := []int{}
		for _, s := range lStrs {
			digit, isInt := util.TryStringToInt(s)
			if isInt {
				lDigits = append(lDigits, digit)
			}
		}
		digits = append(digits, lDigits)
	}
	fmt.Println(digits)
	nums := []int{}
	for _, lDigits := range digits {
		if len(lDigits) < 1 {
			panic("lDigits is empty")
		}
		firstIndex := 0
		lastIndex := len(lDigits) - 1
		num := util.StringToInt(fmt.Sprintf("%d%d", lDigits[firstIndex], lDigits[lastIndex]))
		nums = append(nums, num)
	}
	fmt.Println("sum:", util.Sum(nums...))
}

var digitWords = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

type validDigit struct {
	index int
	val   int
	str   string
}

func (v validDigit) String() string {
	return fmt.Sprintf("{index: %d, val: %d, raw: %s}", v.index, v.val, v.str)
}

func contains(s []string, search string) bool {
	for _, a := range s {
		if a == search {
			return true
		}
	}
	return false
}

func newValidDigit(raw string, idx int) validDigit {
	i, isInt := util.TryStringToInt(raw)
	if !isInt {
		if !contains(digitWords, raw) {
			panic(fmt.Sprintf("raw: %s not a valid digit word", raw))
		}
		for widx, word := range digitWords {
			if word == raw {
				i = widx
			}
		}
	}
	return validDigit{
		index: idx,
		val:   i,
		str:   raw,
	}
}

func makeEmpty(s string) string {
	r := ""
	for range s {
		r += " "
	}
	return r
}

func extractValidDigits(s string) []validDigit {
	vds := []validDigit{}
	// extract words
	for _, dw := range digitWords {
		howMany := strings.Count(s, dw)
		tempS := s
		for howMany > 0 {
			fmt.Println(s)
			fmt.Println(dw)
			fmt.Println(howMany)
			fmt.Println(tempS)
			fmt.Println(len(tempS))
			// if word is in string, get the index
			idx := strings.Index(tempS, dw)
			vds = append(vds, newValidDigit(dw, idx))
			// rm the dw from s
			tempS = strings.Replace(tempS, dw, makeEmpty(dw), 1)
			howMany -= 1
		}
	}
	// extract nums
	for idx, charStr := range strings.Split(s, "") {
		_, ok := util.TryStringToInt(charStr)
		if ok {
			vds = append(vds, newValidDigit(charStr, idx))
		}
	}
	return vds
}

func Part2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	strs := util.ReaderToStrings(f)
	fmt.Println(strs)
	// filter to just digits
	// validDigits := [][]validDigit{}
	nums := []int{}
	for _, line := range strs {
		fmt.Println(line)
		lineValidDigits := extractValidDigits(line)
		// fmt.Println("unsorted ->", lineValidDigits)
		// sort them
		slices.SortFunc(lineValidDigits, func(a, b validDigit) int {
			return cmp.Compare(a.index, b.index)
		})
		fmt.Println(lineValidDigits)
		// validDigits = append(validDigits, lineValidDigits)
		// sum
		lDigits := lineValidDigits
		if len(lDigits) < 1 {
			panic("lDigits is empty")
		}
		// fmt.Println(lDigits)
		firstIndex := 0
		lastIndex := len(lDigits) - 1
		d0 := lDigits[firstIndex].val
		d1 := lDigits[lastIndex].val
		dStr := fmt.Sprintf("%d%d", d0, d1)
		fmt.Println(dStr)
		num := util.StringToInt(dStr)
		nums = append(nums, num)
	}
	// fmt.Println(validDigits)
	// nums := []int{}
	// for _, lDigits := range validDigits {
	// 	if len(lDigits) < 1 {
	// 		panic("lDigits is empty")
	// 	}
	// 	// fmt.Println(lDigits)
	// 	firstIndex := 0
	// 	lastIndex := len(lDigits) - 1
	// 	d0 := lDigits[firstIndex].val
	// 	d1 := lDigits[lastIndex].val
	// 	dStr := fmt.Sprintf("%d%d", d0, d1)
	// 	fmt.Println(dStr)
	// 	num := util.StringToInt(dStr)
	// 	nums = append(nums, num)
	// }
	fmt.Println("sum:", util.Sum(nums...))
}

func main() {
	// Part2("input_example_part2.txt")
	Part2("input.txt")
	// sample := "16one8sixpmjbvqr1six"
	// fmt.Println(sample)
	// fmt.Println(extractValidDigits(sample))
}
