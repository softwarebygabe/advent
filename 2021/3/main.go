package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type Bits []int

func parseInput(filename string) []Bits {
	result := []Bits{}
	util.EvalEachLine(filename, func(line string) {
		bits := []int{}
		bitStrings := strings.Split(line, "")
		for _, bitString := range bitStrings {
			bit, err := strconv.Atoi(bitString)
			if err != nil {
				panic(err)
			}
			bits = append(bits, bit)
		}
		result = append(result, bits)
	})
	return result
}

type mode string

const (
	modeMost  mode = "most"
	modeLeast mode = "least"
)

func common(bits Bits, m mode) int {
	seenMap := map[int]int{} // bit: count
	for _, b := range bits {
		c, ok := seenMap[b]
		if !ok {
			seenMap[b] = 1
		} else {
			seenMap[b] = c + 1
		}
	}
	switch m {
	case modeMost:
		if seenMap[0] > seenMap[1] {
			return 0
		} else {
			return 1
		}
	case modeLeast:
		if seenMap[0] <= seenMap[1] {
			return 0
		} else {
			return 1
		}
	}
	return 0
}

func getBitsByPosition(bitList []Bits, pos int) Bits {
	result := []int{}
	for _, bits := range bitList {
		result = append(result, bits[pos])
	}
	return result
}

func rateBits(bitList []Bits) (gamma, epsilon Bits) {
	gamma = []int{}
	epsilon = []int{}
	for i := range bitList[0] {
		positionalBits := getBitsByPosition(bitList, i)
		gammaBit := common(positionalBits, modeMost)
		epsilonBit := common(positionalBits, modeLeast)
		gamma = append(gamma, gammaBit)
		epsilon = append(epsilon, epsilonBit)
	}
	return
}

func (b Bits) String() string {
	s := ""
	for _, bit := range b {
		s += strconv.Itoa(bit)
	}
	return s
}

func (b Bits) Int() int {
	res, err := strconv.ParseInt(b.String(), 2, 64)
	if err != nil {
		panic(err)
	}
	return int(res)
}

func filterByPosBit(bitList []Bits, pos, b int) []Bits {
	result := []Bits{}
	for _, bits := range bitList {
		if bits[pos] == b {
			result = append(result, bits)
		}
	}
	return result
}

func oxyGenRating(bitList []Bits) Bits {
	pos := 0
	totalLen := len(bitList[0])
	for pos < totalLen {
		if len(bitList) == 1 {
			return bitList[0]
		}
		// find the most common bit at pos
		posBits := getBitsByPosition(bitList, pos)
		commonBit := common(posBits, modeMost)
		// filter bitList by common bit at pos
		remaining := filterByPosBit(bitList, pos, commonBit)
		bitList = remaining
		if len(bitList) == 1 {
			return bitList[0]
		}
		pos++
	}
	return []int{}
}

func co2ScrubRating(bitList []Bits) Bits {
	pos := 0
	totalLen := len(bitList[0])
	for pos < totalLen {
		if len(bitList) == 1 {
			return bitList[0]
		}
		// find the most common bit at pos
		posBits := getBitsByPosition(bitList, pos)
		uncommonBit := common(posBits, modeLeast)
		// filter bitList by common bit at pos
		remaining := filterByPosBit(bitList, pos, uncommonBit)
		bitList = remaining
		if len(bitList) == 1 {
			return bitList[0]
		}
		pos++
	}
	return []int{}
}

func Part1() {
	inputBitList := parseInput("./input.txt")
	gammaB, epsilonB := rateBits(inputBitList)
	fmt.Println(gammaB)
	fmt.Println(epsilonB)
	fmt.Println(gammaB.Int())
	fmt.Println(epsilonB.Int())
	fmt.Println(gammaB.Int() * epsilonB.Int())
}

func main() {
	inputBitList := parseInput("./input.txt")
	oxyGenRateB := oxyGenRating(inputBitList)
	fmt.Println(oxyGenRateB.String())
	fmt.Println(oxyGenRateB.Int())
	co2ScrubRateB := co2ScrubRating(inputBitList)
	fmt.Println(co2ScrubRateB.String())
	fmt.Println(co2ScrubRateB.Int())

	fmt.Println(oxyGenRateB.Int() * co2ScrubRateB.Int())
}
