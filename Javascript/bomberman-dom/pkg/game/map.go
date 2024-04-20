package game

import "fmt"

type Board [13][13]string

func CreateMap() *Board {
	var board = &Board{
		{"wall","wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall"},
		{"wall","empty", "empty", "box", "box", "box", "box", "box", "box", "box", "empty", "empty", "wall"},
		{"wall","empty", "wall", "box", "wall", "box", "wall", "box", "wall", "box", "wall", "empty", "wall"},
		{"wall","box", "box", "box", "box", "box", "box", "box", "box", "box", "box", "box", "wall"},
		{"wall","box", "wall", "box", "wall", "box", "wall", "box", "wall", "box", "wall", "box", "wall"},
		{"wall","box", "box", "box", "box", "box", "box", "box", "box", "box", "box", "box", "wall"},
		{"wall","box", "wall", "box", "wall", "box", "wall", "box", "wall", "box", "wall", "box", "wall"},
		{"wall","box", "box", "box", "box", "box", "box", "box", "box", "box", "box", "box", "wall"},
		{"wall","box", "wall", "box", "wall", "box", "wall", "box", "wall", "box", "wall", "box", "wall"},
		{"wall","box", "box", "box", "box", "box", "box", "box", "box", "box", "box", "box", "wall"},
		{"wall","empty", "wall", "box", "wall", "box", "wall", "box", "wall", "box", "wall", "empty", "wall"},
		{"wall","empty", "empty", "box", "box", "box", "box", "box", "box", "box", "empty", "empty", "wall"},
		{"wall","wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall", "wall"},
	}

	return board
}

func (board *Board) PrintMap() {
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}
}
