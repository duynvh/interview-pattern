package main

import (
	"fmt"
	"math"
	"strings"
)

type ticTacToe struct {
	rows         []int
	cols         []int
	diagonal     int
	antiDiagonal int
}

func (this *ticTacToe) ticTacToeInit(n int) {
	this.rows = make([]int, n)
	this.cols = make([]int, n)

	for i := 0; i < n; i++ {
		this.rows[i] = 0
		this.cols[i] = 0
	}

	this.diagonal = 0
	this.antiDiagonal = 0
}

func (this *ticTacToe) move(row int, col int, player int) int {
	currentPlayer := -1
	if player == 1 {
		currentPlayer = 1
	}

	this.rows[row] += currentPlayer
	this.rows[col] += currentPlayer

	if row == col {
		this.diagonal += currentPlayer
	}

	if col == len(this.cols)-row-1 {
		this.antiDiagonal += currentPlayer
	}

	n := len(this.rows)

	if math.Abs(float64(this.rows[row])) == float64(n) ||
		math.Abs(float64(this.cols[col])) == float64(n) ||
		math.Abs(float64(this.diagonal)) == float64(n) ||
		math.Abs(float64(this.antiDiagonal)) == float64(n) {
		return player
	}
	return 0
}

// Driver code
func main() {
	n := 3
	inputs := [][]int{
		{0, 0, 1},
		{0, 2, 2},
		{2, 2, 1},
		{1, 1, 2},
		{2, 0, 1},
		{1, 0, 2},
		{2, 1, 1},
	}
	ticTacToeObj := new(ticTacToe)
	ticTacToeObj.ticTacToeInit(n)
	win := 0

	for i, v := range inputs {
		fmt.Printf("Move %d.\tPlayer %d makes a move at Move(%d, %d, %d)\n", i, v[2], v[0], v[1], v[2])
		win = ticTacToeObj.move(v[0], v[1], v[2])
		if win == 0 {
			fmt.Printf("\tNo one wins\n")
		} else {
			fmt.Printf("\tPlayer %d wins\n", win)
			break
		}
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
