package main

import (
	"fmt"
	"strings"
)

func findLongestSubstring(inputStr string) int {
	if len(inputStr) == 0 {
		return 0
	}

	n := len(inputStr)
	stCurr, longest, currLen, i := 0, 0, 0, 0
	window := make(map[byte]int)

	for i = 0; i < n; i++ {
		character := inputStr[i]
		if _, ok := window[character]; ok {
			if window[character] >= stCurr {
				currLen = i - stCurr
				if longest < currLen {
					longest = currLen
				}

				stCurr = window[character] + 1
			}
		}

		window[character] = i
	}

	if longest < i-stCurr {
		longest = i - stCurr
	}

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	str := []string{"abcabcbb", "pwwkew", "bbbbb", "ababababa", "", "ABCDEFGHI", "ABCDEDCBA", "AAAABBBBCCCCDDDD"}
	for i, s := range str {
		fmt.Printf("%d.\tInput string: \"%s\"\n", i+1, s)
		result := findLongestSubstring(s)
		fmt.Printf("\tLength of the longest substring: %d\n", result)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
