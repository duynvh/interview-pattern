package main

import (
	"fmt"
	"strings"
)

func verifyAlienDictionary(words []string, order string) bool {
	if len(words) == 1 {
		return true
	}

	orderMap := make(map[byte]int)

	for i := 0; i < len(order); i++ {
		orderMap[order[i]] = i
	}

	// Traverse in array words
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words[i+1]) {
				return false
			}

			if words[i][j] != words[i+1][j] {
				if orderMap[words[i][j]] > orderMap[words[i+1][j]] {
					return false
				} else {
					break
				}
			}
		}
	}

	return true
}

// Driver Code
func main() {
	words := [][]string{
		// {"alpha", "bravo", "charlie", "delta"},
		{"apple", "app"},
		// {"martian"},
		// {"jupyter", "ascending"},
		// {"passengers", "to", "the", "unknown"},
	}
	order := []string{
        "abcdefghijklmnopqrstuvwxyz", 
        // "abcdefghijklmnopqrstuvwxyz", 
        // "mabcdefghijklnopqrstuvwxyz", 
        // "jabcdefghiklmnopqrstuvwxyz", 
        // "ptuhabcdefghijklmnoqrsvwxyz",
    }

	for i, word := range words {
		fmt.Printf("%d.\tWords: ['%s']\n", i+1, strings.Join(word, "', '"))
		fmt.Printf("\tOrder: \"%s\"\n", order[i])
		result := verifyAlienDictionary(word, order[i])
		if result == true {
			fmt.Print("\tAlien dictionary verified: Yes\n")
		} else {
			fmt.Print("\tAlien dictionary verified: No\n")
		}
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
