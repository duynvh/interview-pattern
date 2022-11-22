// https://leetcode.com/problems/minimum-window-subsequence/ (727 - Hard)

// Solution summary
// 1.Initialize two pointers, indexS1 and indexS2, to zero for iterating both strings.
// 2. If the character pointed by indexS1 in str1 is the same as the character pointed by indexS2 in str2, increment both pointers.
// 3. Create two new pointers (start and end) once indexS2 reaches the end of str2. With these two pointers, we’ll slide the window in the opposite direction.
// 4. Set the start to the value of indexS1 and end to start + 1.
// 5. Decrement the start pointer until indexS2 becomes less than zero.
// 6. Decrement indexS2 only if the character pointed by the start pointer in str1 is equal to the character pointed to by indexS2 in str2.
// 7. Calculate the length of a substring by subtracting values of the end and start variables.
// 8. If this length is less than the current minimum length, update the length variable and minSubsequence string.
// 9. Repeat until indexS1 reaches the end of str1.
package main

import (
	"fmt"
	"math"
	"strings"
)

func minWindow(str1 string, str2 string) string {
	sizeStr1, sizeStr2 := len(str1), len(str2)

	length := math.MaxInt64
	_ = length

	indexS1, indexS2 := 0, 0
	minSubsequence := ""
	_ = indexS1
	_ = indexS2
	_ = minSubsequence

	for indexS1 < sizeStr1 {
		if str1[indexS1] == str2[indexS2] {
			indexS2++

			// Check if indexS2 has reached the end of str2
			if indexS2 == sizeStr2 {
				start, end := indexS1, indexS1+1
				_ = start
				_ = end

				// Initialize start to the index where all characters of str2 were present in str1
				indexS2--

				// Decrement pointer indexS2 and start a reverse loop
				for indexS2 >= 0 {
					// Decrement pointer indexS2 until all characters of str2 are found in str1
					if str1[start] == str2[indexS2] {
						indexS2--
					}

					// Decrement start pointer everytime to find the starting point of the required subsequence
					start--
				}

				start++

				// Check if length of subsequence pointed by start and end pointers is less than current min length
				if end-start < length {
					length = end - start
					minSubsequence = str1[start:end]
				}

				// Set indexS1 to start to continue checking in str1 after this discovered subsequence
				indexS1 = start
				indexS2 = 0
			}
		}

		indexS1++
	}

	return minSubsequence
}

// Driver code
func main() {
	str1 := []string{"abcdebdde", "fgrqsqsnodwmxzkzxwqegkndaa", "qwewerrty", "aaabbcbq", "zxcvnhss", "alpha", "beta", "asd", "abcd"}
	str2 := []string{"bde", "kzed", "werty", "abc", "css", "la", "ab", "as", "pp"}

	for index, value := range str1 {
		fmt.Printf("%d.\tInput string: (\"%s\", \"%s\")\n", index+1, value, str2[index])
		fmt.Printf("\tSubsequence string: %s\n", minWindow(value, str2[index]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
