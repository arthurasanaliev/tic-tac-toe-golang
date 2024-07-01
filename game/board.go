package game

import "fmt"

type Board struct {
	grid [3][3]string
}

func NewBoard() *Board {
	return &Board{
		grid: [3][3]string{
			{" ", " ", " "},
			{" ", " ", " "},
			{" ", " ", " "},
		},
	}
}

func (b *Board) Display() {
	for _, row := range b.grid {
		fmt.Println(row)
	}
}

func (b *Board) PlaceMark(x, y int, mark string) bool {
	if x < 0 || x > 2 || y < 0 || y > 2 || b.grid[x][y] != " " {
		return false
	}
	b.grid[x][y] = mark
	return true
}

func (b *Board) CheckWin() bool {
	for i := 0; i < 3; i++ {
		if b.grid[i][0] == b.grid[i][1] && b.grid[i][1] == b.grid[i][2] && b.grid[i][2] != " " {
			return true
		}
	}
	for j := 0; j < 3; j++ {
		if b.grid[0][j] == b.grid[1][j] && b.grid[1][j] == b.grid[2][j] && b.grid[2][j] != " " {
			return true
		}
	}
	if b.grid[0][0] == b.grid[1][1] && b.grid[1][1] == b.grid[2][2] && b.grid[2][2] != " " {
		return true
	}
	if b.grid[0][2] == b.grid[1][1] && b.grid[1][1] == b.grid[2][0] && b.grid[2][0] != " " {
		return true
	}
	return false
}

func (b *Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.grid[i][j] == " " {
				return false
			}
		}
	}
	return true
}
