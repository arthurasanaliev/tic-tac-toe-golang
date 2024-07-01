package game

import "fmt"

type Game struct {
	board         *Board
	player1       *Player
	player2       *Player
	currentPlayer *Player
}

func NewGame() *Game {
	board := NewBoard()
	player1 := NewPlayer("Player 1", "X")
	player2 := NewPlayer("Player 2", "O")

	return &Game{
		board:         board,
		player1:       player1,
		player2:       player2,
		currentPlayer: player1,
	}
}

func (g *Game) Start() {
	for {
		g.board.Display()
		x, y := g.currentPlayer.GetMove()
		if !g.board.PlaceMark(x, y, g.currentPlayer.mark) {
			fmt.Println("Invalid move, try again.")
			continue
		}
		if g.board.CheckWin() {
			g.board.Display()
			fmt.Printf("%s wins!\n", g.currentPlayer.name)
			return
		}
		if g.board.IsFull() {
			g.board.Display()
			fmt.Println("Draw!")
			return
		}
		g.switchPlayer()
	}
}

func (g *Game) switchPlayer() {
	if g.currentPlayer == g.player1 {
		g.currentPlayer = g.player2
	} else {
		g.currentPlayer = g.player1
	}
}
