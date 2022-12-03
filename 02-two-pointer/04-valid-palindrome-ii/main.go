package main

import (
	"fmt"
	"strings"
)

func validPalindrome(s string) bool {
	s = strings.ToLower(s)
	start := 0
	end := len(s) - 1

	for start <= end {
		if s[start] != s[end] {
			if check(s, start+1, end) || check(s, start, end-1) {
				return true
			} else {
				return false
			}
		}

		start++
		end--
	}

	return true
}

func check(s string, start int, end int) bool {
	for start <= end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}

func main() {
	fmt.Println(validPalindrome("abc"))
}
