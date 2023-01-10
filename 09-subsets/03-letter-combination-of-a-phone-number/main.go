package main

import (
	"fmt"
	"strings"
)

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func backtrack(index int, path []string, digits string, letters map[string][]string, combinations []string) []string {
	// If the length of path and digits is same
	// we have a complete combinations
	if len(path) == len(digits) {
		combinations = append(combinations, strings.Join(path, ""))
		return combinations
	}

	// Using the index and digits[index], get the array of letters
	possibleLetters := letters[string(digits[index])]

	for i := 0; i < len(possibleLetters); i++ {
		path = append(path, string(possibleLetters[i]))

		// move on to the next genre
		combinations = backtrack(index+1, path, digits, letters, combinations)

		if len(path) > 0 {
			path = removeIndex(path, len(path)-1)
		}
	}

	return combinations
}

func letterCombinations(digits string) []string {
	// Return empty array if input is empty
	if len(digits) == 0 {
		var temp []string
		return temp
	}

	// Mapping the digits to their corresponding letters
	letters := map[string][]string{
		"1": {""},
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	// Declaring the combinations variable
	var combinations []string

	// Initiate backtracking with an empty path and starting index of 0
	var path []string
	combinations = backtrack(0, path, digits, letters, combinations)
	return combinations
}

func letterCombinations2(digits string) []string {
	// Return empty array if input is empty
	if len(digits) == 0 {
		var temp []string
		return temp
	}

	// Mapping the digits to their corresponding letters
	maps := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	// Declaring the combinations variable
	combinations := []string{""}

	for i := 0; i < len(digits); i++ {
		letters, ok := maps[string(digits[i])]
		if !ok {
			continue
		}
		var tempArr []string
		for j := 0; j < len(letters); j++ {
			letterToAdd := letters[j]
			for k := 0; k < len(combinations); k++ {
				combination := combinations[k]
				tempArr = append(tempArr, fmt.Sprintf("%s%s", combination, letterToAdd))
			}
		}
		combinations = tempArr
	}
	return combinations
}

// Driver code
func main() {
	digits := []string{
		// "2",
		"73",
		// "426",
		// "78",
		// "925",
		// "2345",
	}

	for i, digit := range digits {
		combinations := letterCombinations(digit)
		output := "[\"" + strings.Join(combinations, `", "`) + `"]`
		fmt.Printf("%d.\tAll letter combinations for '%s': %s\n", i+1, digit, output)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
