package main

import (
	"fmt"
	"strings"
)

func permutePalindrome(st string) bool {
	frequencies := make(map[rune]int)
	for _, v := range st {
		present := false
		for j, _ := range frequencies {
			if j == v {
				frequencies[v] += 1
				present = true
			}
		}
		if !present {
			frequencies[v] = 1
		}
	}

	count := 0

	for i, _ := range frequencies {
		if frequencies[i] % 2 != 0 {
			count += 1
		} 
	}

	if count <= 1 {
		return true
	} else {
		return false
	}

	return false
}

// Driver code
func main() {
	strArray := []string{"baefeab", "abc", "xzz", "jjadd", "kllk"}

	for i, v := range strArray {
		fmt.Printf("%d.\tInput string: %s\n", i+1, v)
		result := permutePalindrome(v)
		if result {
			fmt.Printf("\n\tInput string has permutations that are palindromes\n")
		} else {
			fmt.Printf("\n\tInput string does not have permutations that's a palindrome\n")
		}
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
