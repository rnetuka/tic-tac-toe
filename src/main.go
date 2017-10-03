package main

import (
	"flag"
	"strings"
	"strconv"
	"fmt"
	"errors"
)

type Board struct {
	width	int
	height	int
	data	[][]string
}

var difficulty = flag.String("d", "easy", "")
var boardSize = flag.String("s", "3x3", "")
var playerMark = flag.String("m", "x", "")
var playerOrder = flag.Int("o", 1, "")

func main() {
	flag.Parse()

	size := strings.Split(*boardSize, "x")

	width, err := strconv.Atoi(size[0])

	if err != nil {

	}

	height, err := strconv.Atoi(size[1])

	if err != nil {

	}

	game := NewGame(width, height)
	game.board.Mark(0, 0, *game.players[0])

	fmt.Println(*game.board)
}

func NewBoard(width int, height int) *Board {
	board := new(Board)
	board.width = width
	board.height = height
	for i := 0; i< width; i++ {
		column := make([]string, height)

		for j := 0; j < height; j++ {
			column[j] = " "
		}

		board.data = append(board.data, column)
	}
	return board
}

func (board Board) String() string {
	var str string = ""
	for y := 0; y < board.height; y++ {
		for x := 0; x < board.width; x++ {
			str += board.data[x][y]

			if x < board.width - 1 {
				str += "|"
			}
		}
		str += "\n"

		if y < board.height - 1 {

			for x := 0; x < board.width; x++ {
				str += "-"

				if x < board.width-1 {
					str += "+"
				}
			}
			str += "\n"
		}
	}
	return str
}

func (board Board) IsEmpty(x int, y int) (bool, error) {
	if x < 0 || x > board.width {
		return false, errors.New("X must be between 0 and width")
	}
	if y < 0 || y > board.height {
		return false, errors.New("Y must be between 0 and height")
	}
	return board.data[x][y] == " ", nil
}

func (board Board) Mark(x int, y int, player Player) error {
	empty, err := board.IsEmpty(x, y)

	if err != nil {
		return err
	}

	if !empty {
		return errors.New("Coordinates are not empty")
	}
	board.data[x][y] = player.mark
	return nil
}