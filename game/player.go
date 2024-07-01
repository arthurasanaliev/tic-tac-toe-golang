package game

import "fmt"

type Player struct {
	name string
	mark string
}

func NewPlayer(name, mark string) *Player {
	return &Player{name: name, mark: mark}
}

func (p *Player) GetMove() (int, int) {
	var x, y int
	fmt.Printf("%s, enter your move (row and column): ", p.name)
	fmt.Scanf("%d %d", &x, &y)
	return x, y
}
