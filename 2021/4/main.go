package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/softwarebygabe/advent/pkg/util"
)

type Board struct {
	squares       [][]Square
	lastMarkedNum int
}

type Square struct {
	Num      int
	IsMarked bool
}

func (b Board) IsZero() bool {
	return b.squares == nil
}

func (b Board) String() string {
	result := ""
	for _, row := range b.squares {
		for _, sq := range row {
			if sq.IsMarked {
				result += fmt.Sprintf(" [%d]", sq.Num)
			} else {
				result += fmt.Sprintf(" %d", sq.Num)
			}
		}
		result += "\n"
	}
	return result
}

func (b *Board) AddRow(numbers []int) {
	row := []Square{}
	for _, num := range numbers {
		row = append(row, Square{Num: num})
	}
	b.squares = append(b.squares, row)
}

func (b *Board) MarkNum(num int) {
	newSquares := [][]Square{}
	for _, row := range b.squares {
		newRow := []Square{}
		for _, sq := range row {
			if sq.Num == num {
				sq.IsMarked = true
				b.lastMarkedNum = num
			}
			newRow = append(newRow, sq)
		}
		newSquares = append(newSquares, newRow)
	}
	b.squares = newSquares
}

func allMarked(squares []Square) bool {
	isAllMarked := false
	for _, sq := range squares {
		isAllMarked = sq.IsMarked
		if !sq.IsMarked {
			break
		}
	}
	return isAllMarked
}

func (b Board) Won() bool {
	// check all rows
	for _, row := range b.squares {
		rowWon := allMarked(row)
		if rowWon {
			return true
		}
	}
	// check all cols
	for i := range b.squares[0] {
		col := []Square{}
		for _, row := range b.squares {
			col = append(col, row[i])
		}
		colWon := allMarked(col)
		if colWon {
			return true
		}
	}
	return false
}

func (b Board) UnmarkedSum() int {
	var sum int
	for _, row := range b.squares {
		for _, sq := range row {
			if !sq.IsMarked {
				sum += sq.Num
			}
		}
	}
	return sum
}

func (b Board) Score() int {
	return b.lastMarkedNum * b.UnmarkedSum()
}

func PlayBingo(boards []Board, numbers []int) int {
	boardPs := make([]*Board, 0)
	for i := 0; i < len(boards); i++ {
		boardPs = append(boardPs, &boards[i])
	}
	winningBoards := []*Board{}
	remainingBoards := make([]*Board, len(boardPs))
	copy(remainingBoards, boardPs)
	for _, draw := range numbers {
		// fmt.Println("drawing...", draw)
		for i := 0; i < len(remainingBoards); i++ {
			board := remainingBoards[i]
			if board != nil {
				board.MarkNum(draw)
				if board.Won() {
					fmt.Println("bingo! board", i+1)
					winningBoards = append(winningBoards, board)
					remainingBoards[i] = nil
				}
			}
		}
	}
	return winningBoards[len(winningBoards)-1].Score()
}

func parseInput(filepath string) ([]int, []Board) {
	numbers := []int{}
	boards := []Board{}
	lineNum := 0
	newBoard := Board{}
	util.EvalEachLine(filepath, func(line string) {
		if lineNum == 0 {
			splitlist := strings.Split(line, ",")
			for _, substring := range splitlist {
				num, err := strconv.Atoi(substring)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, num)
			}
		} else if lineNum > 1 {
			if len(line) == 0 {
				// newBoard has been populated so add to list
				boards = append(boards, newBoard)
				// reset
				newBoard = Board{}
			} else {
				// non-empty line
				splitlist := strings.Split(line, " ")
				filtered := []string{}
				for _, substring := range splitlist {
					if substring != "" {
						filtered = append(filtered, substring)
					}
				}
				boardNums := []int{}
				for _, substring := range filtered {
					boardNum, err := strconv.Atoi(substring)
					if err != nil {
						panic(err)
					}
					boardNums = append(boardNums, boardNum)
				}
				newBoard.AddRow(boardNums)
			}
		}
		lineNum++
	})
	boards = append(boards, newBoard)
	return numbers, boards
}

func main() {
	numbers, boards := parseInput("./input.txt")
	// fmt.Println(numbers)
	// for _, board := range boards {
	// 	fmt.Println(board.String())
	// }
	fmt.Println("number of boards:", len(boards))
	winningScore := PlayBingo(boards, numbers)
	fmt.Println(winningScore)
}
