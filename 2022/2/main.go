package main

import (
	"fmt"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type moveType int

const (
	rock moveType = iota + 1
	paper
	scissors
)

var winMap = map[moveType]moveType{
	rock:     scissors,
	paper:    rock,
	scissors: paper,
}

var loseMap = map[moveType]moveType{
	scissors: rock,
	rock:     paper,
	paper:    scissors,
}

const (
	loss = 0
	draw = 3
	win  = 6
)

type move struct {
	name moveType
}

func newMove(v string) move {
	switch v {
	case "A", "X":
		return move{rock}
	case "B", "Y":
		return move{paper}
	case "C", "Z":
		return move{scissors}
	default:
		panic("unable to determine move from input")
	}
}

func newMoveFromResult(result string, oppMove moveType) move {
	switch result {
	case "X":
		// lose
		return move{winMap[oppMove]}
	case "Y":
		// draw
		return move{oppMove}
	case "Z":
		// win
		return move{loseMap[oppMove]}
	default:
		panic("unable to determine move from input")
	}
}

func (m move) pts() int {
	return int(m.name)
}

func (m move) resultPts(m2 move) int {
	// check for draw
	if m.name == m2.name {
		return draw
	}
	winsAgainst := winMap[m.name]
	if m2.name == winsAgainst {
		return win
	}
	return loss
}

func parseInput(filename string) [][]string {
	allMoves := [][]string{}
	util.EvalEachLine(filename, func(line string) {
		moves := strings.Split(line, " ")
		allMoves = append(allMoves, moves)
	})
	return allMoves
}

func Part1() {
	allMoves := parseInput("input.txt")
	var totalScore int
	for _, moves := range allMoves {
		oppMove := newMove(moves[0])
		yourMove := newMove(moves[1])
		roundScore := yourMove.pts() + yourMove.resultPts(oppMove)
		// fmt.Println("round score:", roundScore)
		totalScore += roundScore
	}
	fmt.Println("score across all rounds following strategy:", totalScore)
}

func Part2() {
	allMoves := parseInput("input.txt")
	var totalScore int
	for _, moves := range allMoves {
		oppMove := newMove(moves[0])
		yourMove := newMoveFromResult(moves[1], oppMove.name)
		roundScore := yourMove.pts() + yourMove.resultPts(oppMove)
		// fmt.Println("round score:", roundScore)
		totalScore += roundScore
	}
	fmt.Println("score across all rounds following strategy:", totalScore)
}

func main() {
	Part2()
}
