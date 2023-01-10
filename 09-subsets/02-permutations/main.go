package main

import (
	"fmt"
	"strings"
)

func swapChar(word string, i int, j int) string {
	swapIndex := []rune(word)
	swapIndex[i], swapIndex[j] = swapIndex[j], swapIndex[i]
	return string(swapIndex)
}

func permuteStringRec(word string, currentIndex int, result *[]string) {
	// Prevent adding duplicate permutations
	if currentIndex == len(word)-1 {
		*result = append(*result, word)
        return
	}

	for i := currentIndex; i < len(word); i++ {
		swappedStr := swapChar(word, currentIndex, i)

		permuteStringRec(swappedStr, currentIndex+1, result)
	}
}

func permuteWord(word string) []string {
	result := make([]string, 0)

	permuteStringRec(word, 0, &result)

	return result
}

// Driver code
func main() {
	inputList := []string{"ab", "bad", "abcd"}

	for index, value := range inputList {
		fmt.Printf("%d.\tInput string: '%s'\n", index+1, value)
		fmt.Printf("\tAll possible permutations are: [%s]\n", strings.Join(permuteWord(value), ", "))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
