// H(ACG)=H(A)+H(C)+H(G)
// H(CGT)=((H(ACG)−H(A))×4)+H(T)
// https://leetcode.com/problems/repeated-dna-sequences/ 187 - Medium

// Solution summary
// 1. Iterate the string.
// 2. Compute the hash value for the contents of the window.
// 3. Add this hash value to the hash map (being used to perform set functionality) that keeps track of the hashes of all substrings of the given length.
// 4. Move the window one step forward and compute a new hash value.
// 5. If the hash value of a window has already been seen, the sequence in this window is repeated, so we add it to our output map.
// 6. Once all characters of the string have been traversed, we convert the output map to a string array and return the string array.
package main

import (
	"fmt"
	"math"
	"strings"
)

func findRepeatedSequences(s string, k int) []string {
	windowSize := k

	if len(s) <= windowSize {
		return make([]string, 0)
	}

	// Parameters of rolling hash
	base := 4 // 'a', the hash parameter

	// Mapping of a character into an integer
	mapping := map[rune]int{'A': 1, 'C': 2, 'G': 3, 'T': 4}
	numbers := make([]int, 0)
	for _, ch := range s {
		numbers = append(numbers, mapping[ch])
	}

	hashing := 0
	substringHashes, output := make(map[int]bool, 0), make(map[string]int, 0)

	// Iterate over all window-sized sub-strings
	for start := 0; start < len(s)-windowSize+1; start++ {
		// Hash function of current subsequence
		if start != 0 {
			hashing = (hashing - (int(math.Pow(float64(base), float64(windowSize-1))) * numbers[start-1])) * base
			hashing = hashing + numbers[start+windowSize-1]
		} else {
			for end := 0; end < windowSize; end++ {
				hashing += int(math.Pow(float64(base), float64(windowSize-end-1))) * numbers[end]
			}
		}

		// Subsequence and output sets
		if _, ok := substringHashes[hashing]; ok {
			output[s[start:start+windowSize]] = 0
		}
		substringHashes[hashing] = true
	}

	stringOutput := make([]string, 0)
	for str, _ := range output {
		stringOutput = append(stringOutput, str)
	}
	return stringOutput
}

func main() {
	inputsString := []string{"ACGT", "AGACCTAGAC", "AAAAACCCCCAAAAACCCCCC", "GGGGGGGGGGGGGGGGGGGGGGGGG", "TTTTTCCCCCCCTTTTTTCCCCCCCTTTTTTT", "TTTTTGGGTTTTCCA", "", "AAAAAACCCCCCCAAAAAAAACCCCCCCTG", "ATATATATATATATAT"}
	inputsK := []int{3, 3, 8, 12, 10, 14, 10, 10, 6}

	// inputsString := []string{"AGACCTAGAC"}
	// inputsK := []int{3}

	for i := range inputsK {
		fmt.Printf("%d.\tInput Sequence: \"%s\"\n", i+1, inputsString[i])
		fmt.Printf("\tk: %d\n", inputsK[i])
		result := findRepeatedSequences(inputsString[i], inputsK[i])
		fmt.Printf("\tConverted Subsequence: %s\n", strings.Replace(fmt.Sprint(result), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
