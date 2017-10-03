package main

type Game struct {
	board 	*Board
	players [2]*Player
}

func NewGame(width int, height int) *Game {
	game := new(Game)
	game.board = NewBoard(width, height)
	game.players[0] = NewPlayer("x")
	game.players[1] = NewPlayer("o")
	return game
}