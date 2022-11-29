package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {
	left := 0
	right := len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return false
		}

		left++
		right--
	}

	return true
}

// Driver code
func main() {
	str := []string{"RACEACAR", "A", "ABCDEFGFEDCBA", "ABC", "ABCBA", "ABBA", "RACEACAR"}
	for i, s := range str {
		fmt.Printf("Test Case # %d\n", i+1)
		fmt.Printf("\tInput string: \"%s\"\n", s)
		fmt.Printf("\nIs it a palindrome?.....%v\n", isPalindrome(s))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
