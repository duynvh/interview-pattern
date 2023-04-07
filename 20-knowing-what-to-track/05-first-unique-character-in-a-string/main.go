package main

import (
	"fmt"
	"strings"
)

func firstUniqueChar(s string) int {
	characterCountMap := make(map[rune]int)

	for _, chr := range s {
		if _, ok := characterCountMap[chr]; ok {
			characterCountMap[chr]++
		} else {
			characterCountMap[chr] = 1
		}
	}

	for i, chr := range s {
		if characterCountMap[chr] == 1 {
			return i
		}
	}

	return -1
}

// main is used for driver code
func main() {
	strArray := []string{"baefeab", "aabbcc", "dajhfiuebdafsdhdgaj",
		"xyurtwxwtryua", "aeiouqwertyauieotweryqq", "awsjuhfajwfnkag"}

	for i, str := range strArray {
		fmt.Printf("%d.\t Input string: \"%s%s\n", i+1, str, "\"")
		result := firstUniqueChar(str)
		fmt.Print("\t Finding a unique character .....\n")
		fmt.Printf("\t Index of the first unique character is: %d\n", result)
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}
}
