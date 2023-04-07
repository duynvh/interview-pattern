package main

import (
	"fmt"
	"strings"
)

func findAnagrams(s string, p string) []int {
	windowStart, matched := 0, 0
	charFreq := make(map[byte]int)
	resultIndex := make([]int, 0)

	for i := 0; i < len(p); i++ {
		charFreq[p[i]]++
	}

	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		endChar := s[windowEnd]
		if _, ok := charFreq[endChar]; ok {
			charFreq[endChar]--
			if charFreq[endChar] == 0 {
				matched++
			}
		}

		if matched == len(charFreq) {
			resultIndex = append(resultIndex, windowStart)
		}

		if windowEnd >= len(p)-1 {
			startChar := s[windowStart]
			windowStart++

			if _, ok := charFreq[startChar]; ok {
				if charFreq[startChar] == 0 {
					matched--
				}

				charFreq[startChar]++
			}
		}
	}

	return resultIndex
}

func main() {
	A := []string{"abab", "cbaebabacd", "cefecf", "hello", "bro"}
	B := []string{"ab", "abc", "efc", "olleh", "bro"}
	for i := range A {
		fmt.Printf("%d.\tString a: \"%s\"\n", i+1, A[i])
		fmt.Printf("\tString b: \"%s\"\n", B[i])
		output := findAnagrams(A[i], B[i])
		fmt.Printf("\tAnagrams of string b start at index(es) %s in string a.\n", strings.Replace(fmt.Sprint(output), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}
}
