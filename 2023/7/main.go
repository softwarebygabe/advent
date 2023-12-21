package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

const cardRanks = "23456789TJQKA"

var handRanks = []string{
	"high-card",
	"one-pair",
	"two-pair",
	"three-kind",
	"full-house",
	"four-kind",
	"five-kind",
}

type card struct {
	r rune
}

func (c card) String() string {
	return string([]rune{c.r})
}

func (c card) rank() int {
	return strings.IndexRune(cardRanks, c.r)
}

func (c card) isEqual(c2 card) bool {
	return c.rank() == c2.rank()
}

type hand struct {
	cards []card
	rank  string
	bid   int
}

func (h hand) String() string {
	return fmt.Sprintf("%v %10s bid: %d", h.cards, h.rank, h.bid)
}

func newHand(bid int, cards ...card) hand {
	h := hand{
		cards: cards,
		bid:   bid,
	}
	cardCounts := map[rune]int{}
	for _, card := range cards {
		count, ok := cardCounts[card.r]
		if !ok {
			cardCounts[card.r] = 1
		} else {
			cardCounts[card.r] = count + 1
		}
	}
	counter := util.NewMinerMaxer()
	for _, v := range cardCounts {
		counter.Add(v)
	}
	if counter.GetMax() == 5 {
		h.rank = "five-kind"
		return h
	}
	if counter.GetMax() == 4 {
		h.rank = "four-kind"
		return h
	}
	if counter.GetMax() == 3 {
		if counter.GetMin() == 2 {
			h.rank = "full-house"
			return h
		}
		h.rank = "three-kind"
		return h
	}
	if counter.GetMax() == 2 {
		pairs := 0
		for _, v := range cardCounts {
			if v > 1 {
				pairs++
			}
		}
		if pairs > 1 {
			h.rank = "two-pair"
			return h
		}
		h.rank = "one-pair"
		return h
	}
	h.rank = "high-card"
	return h
}

func (h hand) rankNum() int {
	for i, hr := range handRanks {
		if h.rank == hr {
			return i
		}
	}
	panic("rank num unknown")
}

func (h hand) lessThan(h2 hand) bool {
	if h.rankNum() == h2.rankNum() {
		for i := 0; i < len(h.cards); i++ {
			c1 := h.cards[i]
			c2 := h2.cards[i]
			if c1.isEqual(c2) {
				continue
			}
			return c1.rank() < c2.rank()
		}
	}
	return h.rankNum() < h2.rankNum()
}

type byRank []hand

func (h byRank) Len() int {
	return len(h)
}

func (h byRank) Less(i, j int) bool {
	return h[i].lessThan(h[j])
}

func (h byRank) Swap(i, j int) {
	temp := h[i]
	h[i] = h[j]
	h[j] = temp
}

func part1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	lines := util.ReaderToStrings(f)
	hands := []hand{}
	for _, line := range lines {
		cardsRaw, bidRaw, ok := strings.Cut(line, " ")
		if !ok {
			panic("could not parse line")
		}
		cards := []card{}
		for _, r := range cardsRaw {
			cards = append(cards, card{r})
		}
		bid, err := strconv.Atoi(bidRaw)
		if err != nil {
			panic("could not parse bid")
		}
		hands = append(hands, newHand(bid, cards...))
	}
	for _, h := range hands {
		fmt.Println(h)
	}
	sort.Sort(byRank(hands))
	fmt.Println("--- sorted ---")
	winnings := []int{}
	for idx, h := range hands {
		fmt.Println(h)
		winnings = append(winnings, h.bid*(idx+1))
	}
	fmt.Println("sum:", util.Sum(winnings...))
}

func main() {
	part1("input.txt")
}
