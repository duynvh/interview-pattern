package main

import (
	"fmt"
	"strings"
)

func isIsomorphic(str1 string, str2 string) bool {
	mapStr1Str2 := make(map[string]string)
	mapStr2Str1 := make(map[string]string)

	for index, _ := range str1 {
		char1 := string(str1[index])
		char2 := string(str2[index])

		if _, ok := mapStr1Str2[char1]; ok && mapStr1Str2[char1] != char2 {
			return false
		}

		if _, ok := mapStr2Str1[char2]; ok && mapStr2Str1[char2] != char1 {
			return false
		}

		mapStr1Str2[char1] = char2
		mapStr2Str1[char2] = char1
	}

	return true
}

// Driver code
func main() {
	A := []string{"egg", "foo", "paper", "badc", "aaeaa"}
	B := []string{"all", "bar", "title", "baba", "uuxyy"}

	for index, value := range A {
		fmt.Printf("%d.\tString 1: \"%s\"\n", index+1, value)
		fmt.Printf("\tString 2: \"%s\"\n", B[index])
		fmt.Printf("\n\tIsomorphic String ? %t\n", isIsomorphic(value, B[index]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
