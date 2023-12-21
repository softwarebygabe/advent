package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/softwarebygabe/advent/pkg/util"
)

type card struct {
	winning map[string]bool
	nums    []string
}

func newCard(w, n []string) card {
	c := card{
		winning: make(map[string]bool),
		nums:    n,
	}
	for _, wn := range w {
		c.winning[wn] = true
	}
	return c
}

func (c card) points() int {
	total := 0
	for _, n := range c.nums {
		if _, ok := c.winning[n]; ok {
			// winning number!
			if total == 0 {
				total = 1 // one point
			} else {
				total *= 2 // double it
			}
		}
	}
	return total
}

func (c card) winningCount() int {
	count := 0
	for _, n := range c.nums {
		if _, ok := c.winning[n]; ok {
			// winning number!
			count += 1
		}
	}
	return count
}

func trimAllSpace(s string) string {
	res := []rune{}
	for _, r := range s {
		if !unicode.IsSpace(r) {
			res = append(res, r)
		}
	}
	return string(res)
}

func cleanEach(s []string) []string {
	res := []string{}
	for _, si := range s {
		if si != "" {
			res = append(res, trimAllSpace(si))
		}
	}
	return res
}

func part1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)

	// parse cards
	cards := []card{}
	for i, line := range lines {
		line = strings.TrimPrefix(line, fmt.Sprintf("Card %d: ", i+1))
		winningRaw, numsRaw, ok := strings.Cut(line, " | ")
		if !ok {
			panic("could not cut on `|`")
		}
		winning := cleanEach(strings.Split(winningRaw, " "))
		nums := cleanEach(strings.Split(numsRaw, " "))
		cards = append(cards, newCard(winning, nums))
	}

	// count up points
	total := 0
	for _, c := range cards {
		fmt.Println("card", c)
		fmt.Println("points", c.points())
		total += c.points()
	}
	fmt.Println("total points:", total)
}

func part2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)

	// parse cards
	cards := []card{}
	for i, line := range lines {
		line = strings.TrimPrefix(line, fmt.Sprintf("Card %d: ", i+1))
		winningRaw, numsRaw, ok := strings.Cut(line, " | ")
		if !ok {
			panic("could not cut on `|`")
		}
		winning := cleanEach(strings.Split(winningRaw, " "))
		nums := cleanEach(strings.Split(numsRaw, " "))
		cards = append(cards, newCard(winning, nums))
	}
	cardCopyTracker := make([]int, len(cards))
	for idx, card := range cards {
		fmt.Println("card", idx+1, "winning count:", card.winningCount())
		for w := 0; w < card.winningCount(); w++ {
			copiedCardIdx := idx + 1 + w // next one and move with w
			cardCopyTracker[copiedCardIdx] = cardCopyTracker[copiedCardIdx] + 1
		}
		if idx > 0 {
			// process copies as well
			fmt.Printf("processing %d copies of card %d\n", cardCopyTracker[idx], idx+1)
			copyCount := cardCopyTracker[idx]
			for copyCount > 0 {
				for w := 0; w < card.winningCount(); w++ {
					copiedCardIdx := idx + 1 + w // next one and move with w
					cardCopyTracker[copiedCardIdx] = cardCopyTracker[copiedCardIdx] + 1
				}
				copyCount--
			}
		}
	}
	fmt.Println(cardCopyTracker)
	fmt.Println("copies:", util.Sum(cardCopyTracker...))
	fmt.Println("orig:", len(cards))
	fmt.Println("total cards:", util.Sum(cardCopyTracker...)+len(cards))
}

func main() {
	part2("./input.txt")
}
