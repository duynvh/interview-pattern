package main

import (
	"fmt"
	"strings"
)

// func wordSearch(grid [][]string, word string) bool {
// 	n := len(grid)
// 	if n < 1 {
// 		return false
// 	}

// 	m := len(grid[0])
// 	if m < 1 {
// 		return false
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			if depthFirstSearch(i, j, word, grid) {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }

// func depthFirstSearch(row int, col int, word string, grid [][]string) bool {
// 	if len(word) == 0 {
// 		return true
// 	}

// 	if row < 0 || row == len(grid) || col < 0 || col == len(grid[0]) || strings.ToLower(grid[row][col]) != strings.ToLower(string(word[0])) {
// 		return false
// 	}

// 	result := false
// 	grid[row][col] = "*"

// 	for _, v := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
// 		result = depthFirstSearch(row+v[0], col+v[1], word[1:], grid)

// 		if result {
// 			break
// 		}
// 	}

// 	grid[row][col] = string(word[0])
// 	return result
// }

func wordSearch(board [][]byte, word string) bool {
	n := len(board)
	if n < 1 {
		return false
	}

	m := len(board[0])
	if m < 1 {
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if depthFirstSearch(i, j, word, board) {
				return true
			}
		}
	}

	return false
}

func depthFirstSearch(row int, col int, word string, board [][]byte) bool {
	if len(word) == 0 {
		return true
	}

	if row < 0 || row == len(board) || col < 0 || col == len(board[0]) || board[row][col] != word[0] {
		return false
	}

	result := false
	board[row][col] = '0'

	for _, v := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		result = depthFirstSearch(row+v[0], col+v[1], word[1:], board)
		if result {
			break
		}
	}

	board[row][col] = word[0]
	return result
}

// Driver code
func main() {
	inputListByte := [][][]byte{
		{
			{'E', 'D', 'X', 'I', 'W'},
			{'P', 'U', 'F', 'M', 'Q'},
			{'I', 'C', 'Q', 'R', 'F'},
			{'M', 'A', 'L', 'C', 'A'},
			{'J', 'T', 'I', 'V', 'E'},
		},
	}

	// inputList := [][][]string{
	// 	{
	// 		{"E", "D", "X", "I", "W"},
	// 		{"P", "U", "F", "M", "Q"},
	// 		{"I", "C", "Q", "R", "F"},
	// 		{"M", "A", "L", "C", "A"},
	// 		{"J", "T", "I", "V", "E"},
	// 	},
	// 	// {
	// 	// 	{"O", "Y", "O", "I"},
	// 	// 	{"B", "I", "E", "M"},
	// 	// 	{"K", "D", "Y", "R"},
	// 	// 	{"M", "T", "W", "I"},
	// 	// 	{"Z", "I", "T", "O"},
	// 	// },
	// 	// {
	// 	// 	{"h", "e", "c", "m", "l"},
	// 	// 	{"w", "l", "i", "e", "u"},
	// 	// 	{"a", "r", "r", "s", "n"},
	// 	// 	{"s", "i", "i", "o", "r"},
	// 	// },
	// 	// {
	// 	// 	{"C", "Q", "N", "A"},
	// 	// 	{"P", "S", "E", "I"},
	// 	// 	{"Z", "A", "P", "E"},
	// 	// 	{"J", "V", "T", "K"},
	// 	// },
	// 	// {
	// 	// 	{"A"},
	// 	// },
	// 	// {
	// 	// 	{"T", "P", "S", "I", "W", "P", "S", "G"},
	// 	// 	{"P", "I", "Y", "C", "Q", "T", "A", "Y"},
	// 	// 	{"I", "S", "L", "H", "F", "Y", "D", "Q"},
	// 	// 	{"M", "T", "A", "I", "N", "H", "X", "E"},
	// 	// 	{"J", "I", "N", "A", "R", "K", "R", "T"},
	// 	// 	{"I", "M", "D", "T", "R", "F", "B", "N"},
	// 	// 	{"I", "M", "D", "Z", "I", "F", "M", "J"},
	// 	// 	{"I", "M", "D", "Z", "S", "T", "W", "I"},
	// 	// },
	// }
	wordsList := []string{
		"EDUCATIVE",
		// "DYNAMIC",
		// "WARRIOR",
		// "save",
		// "ABC",
		// "PSYCHIATRIST",
	}

	for index, value := range inputListByte {
		searchResult := wordSearch(value, wordsList[index])
		if searchResult {
			fmt.Printf("\n\tSearch result = Found word\n")
		} else {
			fmt.Printf("\n\tSearch result = Word not found\n")
		}
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
