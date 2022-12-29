package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func reorganizeString(inputString string) string {
	// Find maximum frequency element from string
	charCounter := make(map[rune]int)
	for _, c := range inputString {
		freq := 1
		if _, ok := charCounter[c]; ok {
			freq = charCounter[c] + 1
		}

		charCounter[c] = freq
	}

	// Store counts of each char in the string along with their occurrence
	mostFreqChars := newMaxHeap()
	for i, _ := range charCounter {
		heap.Push(mostFreqChars, Set{char: i, count: charCounter[i]})
	}

	// Initialzing variables
	result := ""
	previous := new(Set)

	// Store the frequencies of element in max heap structure
	for !mostFreqChars.Empty() || (previous.count != 0 && previous.char != 0) {
		if (previous.count != 0 && previous.char != 0) && mostFreqChars.Empty() {
			return ""
		}

		poppedElement := heap.Pop(mostFreqChars).(Set)
		count, char := poppedElement.count, poppedElement.char
		result += string(char)
		count -= 1

		// Pushing the char back to heap
		if previous.count != 0 && previous.char != 0 {
			heap.Push(mostFreqChars, *previous)
			previous = new(Set)
		}

		if count != 0 {
			previous.count = count
			previous.char = char
		}
	}

	return result + ""
}

// Driver code
func main() {
	testCases := []string{
		// "programming",
		// "hello",
		// "fofjjb",
		// "abbacdde",
		// "aba",
		// "awesome",
		"aaab",
	}

	for i, s := range testCases {
		fmt.Printf("%d.\tInput String: \"%s\"\n", i+1, s)
		fmt.Printf("\tReorganized string: \"%s\"\n", reorganizeString(s))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
