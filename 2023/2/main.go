package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

/*
which games would have been possible if the bag contained only
12 red cubes, 13 green cubes, and 14 blue cubes?
*/

type game struct {
	id    int
	red   *util.MinerMaxer
	green *util.MinerMaxer
	blue  *util.MinerMaxer
}

func newGame(id int) game {
	return game{
		id:    id,
		red:   util.NewMinerMaxer(),
		green: util.NewMinerMaxer(),
		blue:  util.NewMinerMaxer(),
	}
}

func (g game) String() string {
	return fmt.Sprintf("game: %d\nmaxes: %d (r) %d (g) %d (b)",
		g.id,
		g.red.GetMax(),
		g.green.GetMax(),
		g.blue.GetMax(),
	)
}

func (g game) IsPossible(r, grn, b int) bool {
	rmax := g.red.GetMax()
	gmax := g.green.GetMax()
	bmax := g.blue.GetMax()
	return rmax <= r && gmax <= grn && bmax <= b
}

func (g game) Power() int {
	rmax := g.red.GetMax()
	gmax := g.green.GetMax()
	bmax := g.blue.GetMax()
	return rmax * gmax * bmax
}

func parseGame(line string) game {
	gameIDRaw, gameRollsRaw, ok := strings.Cut(line, ": ")
	if !ok {
		panic("colon not found")
	}
	gameID, err := strconv.Atoi(strings.Split(gameIDRaw, " ")[1])
	if err != nil {
		panic("could not parse game id")
	}
	g := newGame(gameID)
	// go through game rolls and add info to game
	gameRolls := strings.Split(gameRollsRaw, "; ")
	for _, gameRoll := range gameRolls {
		diceShown := strings.Split(gameRoll, ", ")
		for _, die := range diceShown {
			numRaw, color, ok := strings.Cut(die, " ")
			if !ok {
				panic("could not parse die")
			}
			num, err := strconv.Atoi(numRaw)
			if err != nil {
				panic("could not parse num die")
			}
			switch color {
			case "red":
				g.red.Add(num)
			case "green":
				g.green.Add(num)
			case "blue":
				g.blue.Add(num)
			}
		}
	}
	return g
}

func part1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	strs := util.ReaderToStrings(f)
	// fmt.Println(strs)
	// parse to games
	games := []game{}
	for _, line := range strs {
		fmt.Println(line)
		g := parseGame(line)
		games = append(games, g)
	}
	possibleGameIDs := []int{}
	for _, game := range games {
		fmt.Println(game)
		if game.IsPossible(12, 13, 14) {
			possibleGameIDs = append(possibleGameIDs, game.id)
		}
	}
	fmt.Println("possible games -->", possibleGameIDs)
	fmt.Println("sum:", util.Sum(possibleGameIDs...))
}

func part2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	strs := util.ReaderToStrings(f)
	// fmt.Println(strs)
	// parse to games
	games := []game{}
	for _, line := range strs {
		fmt.Println(line)
		g := parseGame(line)
		games = append(games, g)
	}
	gamePowers := []int{}
	for _, game := range games {
		fmt.Println(game)
		gamePowers = append(gamePowers, game.Power())
	}
	fmt.Println("games powers-->", gamePowers)
	fmt.Println("sum:", util.Sum(gamePowers...))
}

func main() {
	part2("input.txt")
}
