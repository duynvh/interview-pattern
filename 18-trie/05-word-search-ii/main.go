package main

import (
	"fmt"
	"strings"
)

func dfs(node *TrieNode, grid [][]rune, row int, col int, result *[]string, word string) {
	if node.isString {
		(*result) = append((*result), word)
		node.isString = false
	}

	if 0 <= row && row < len(grid) && 0 <= col && col < len(grid[0]) {
		c := rune(grid[row][col])

		if _, ok := node.children[c]; ok && c != '#' {
			child := node.children[c]
			word += string(c)

			grid[row][col] = '#'
			offsets := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
			for _, offset := range offsets {
				rowOffset := offset[0]
				colOffset := offset[1]

				dfs(child, grid, row+rowOffset, col+colOffset, result, word)
			}

			grid[row][col] = c
		}
	}
}

func findStrings(grid [][]rune, words []string) []string {
	trieForWords := &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
	result := make([]string, 0)
	for _, word := range words {
		trieForWords.Insert(word)
	}

	word := ""
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			dfs(trieForWords.root, grid, row, col, &result, word)
		}
	}

	return result
}

// printGrid is used for printing purposes
func printGrid(grid [][]rune) {
	for _, v := range grid {
		fmt.Printf("\t")
		for _, v1 := range v {
			fmt.Printf("%s   ", string(v1))
		}
		fmt.Printf("\n")
	}
}

// stringArrayToString is used to convert a string array to string.
// It is used for printing purposes in driver code.
func stringArrayToString(str []string) string {
	res := "["
	for j, s := range str {
		res += "\"" + s + "\""
		if j < len(str)-1 {
			res += ", "
		}
	}
	res += "]"
	return res
}

// Driver code
func main() {
	grid1 := [][]rune{{'B', 'S', 'L', 'I', 'M'},
		{'R', 'I', 'L', 'M', 'O'},
		{'O', 'L', 'I', 'E', 'O'},
		{'R', 'Y', 'I', 'L', 'N'},
		{'B', 'U', 'N', 'E', 'C'}}

	grid2 := [][]rune{{'C', 'S', 'L', 'I', 'M'},
		{'O', 'I', 'B', 'M', 'O'},
		{'O', 'L', 'U', 'E', 'O'},
		{'N', 'L', 'Y', 'S', 'N'},
		{'S', 'I', 'N', 'E', 'C'}}

	grid3 := [][]rune{{'C', 'O', 'L', 'I', 'M'},
		{'I', 'N', 'L', 'M', 'O'},
		{'A', 'L', 'I', 'E', 'O'},
		{'R', 'T', 'A', 'S', 'N'},
		{'S', 'I', 'T', 'A', 'C'}}

	grid4 := [][]rune{{'P', 'S', 'L', 'A', 'M'},
		{'O', 'P', 'U', 'R', 'O'},
		{'O', 'L', 'I', 'E', 'O'},
		{'R', 'T', 'A', 'S', 'N'},
		{'S', 'I', 'T', 'A', 'C'}}

	grid5 := [][]rune{{'O', 'A', 'A', 'N'},
		{'E', 'T', 'A', 'E'},
		{'I', 'H', 'K', 'R'},
		{'I', 'F', 'L', 'V'}}

	grid6 := [][]rune{{'S', 'T', 'R', 'A', 'C'},
		{'I', 'R', 'E', 'E', 'E'},
		{'N', 'G', 'I', 'T', 'C'},
		{'I', 'T', 'S', 'R', 'A'}}

	grid7 := [][]rune{{'A', 'A', 'A'},
		{'A', 'A', 'A'},
		{'A', 'A', 'A'}}

	testCaseGrid := [][][]rune{grid1, grid2, grid3, grid4, grid5, grid6, grid7}
	strings1 := []string{"BUY", "SLICK", "SLIME", "ONLINE", "NOW"}
	strings2 := []string{"BUY", "STUFF", "ONLINE", "NOW"}
	strings3 := []string{"REINDEER", "IN", "RAIN"}
	strings4 := []string{"TOURISM", "DESTINY", "POPULAR"}
	strings5 := []string{"OATH", "PEA", "EAT", "RAIN"}
	strings6 := []string{"STREET", "STREETCAR", "STRING", "STING", "RING", "RACECAR"}
	strings7 := []string{"A", "AA", "AAA", "AAAA"}
	stringInput := [][]string{strings1, strings2, strings3, strings4, strings5, strings6, strings7}

	for i := range testCaseGrid {
		if i > 0 {
			fmt.Print("\n")
		}
		fmt.Printf("%d.\t2D Grid:\n\n", i+1)
		printGrid(testCaseGrid[i])
		fmt.Printf("\n\tInput list: %s\n", stringArrayToString(stringInput[i]))
		result := findStrings(testCaseGrid[i], stringInput[i])
		fmt.Printf("\n\tOutput: %s\n", stringArrayToString(result))
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}
}
