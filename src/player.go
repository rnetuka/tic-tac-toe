package main

type Player struct {
	mark string
}

func NewPlayer(mark string) *Player {
	player := new(Player)
	player.mark = mark
	return player
}
