package main

import (
	"fmt"
	"strings"
)

func removeDuplicates2(toCleanUp string) string {
    stack := []rune{}
    for _, char := range toCleanUp {
        if len(stack) == 0 || stack[len(stack) - 1] != char {
            stack = append(stack, char)
        } else {
            stack = stack[:len(stack) - 1]
        }
    }

    return string(stack)
}

func removeDuplicates(toCleanUp string) string {
	frequencyStack := new(Stack)
	k := 2

	for _, char := range toCleanUp {
		if _, ok := frequencyStack.Top()[char]; !frequencyStack.Empty() && ok {
			frequencyStack.Top()[char] += 1
			if frequencyStack.Top()[char] == k {
				frequencyStack.Pop()
			}
		} else {
			frequencyStack.Push(char, 1)
		}
	}

	result := ""
	for _, data := range frequencyStack.s {
		for char, count := range data {
			result += strings.Repeat(string(char), count)
		}
	}

	return result
}

// main is used for driver code
func main() {
	inputs := []string{
		"g",
		"ggaabcdeb",
		"abbddaccaaabcd",
		"aabbccdd",
		"aannkwwwkkkwna",
		"abbabccblkklu",
	}
	for i, input := range inputs {
		fmt.Printf("%d.\tRemove duplicates from string: \"%s\"\n", i+1, input)
		resultingString := removeDuplicates2(input)
		fmt.Printf("\tString after removing duplicates: \"%s\"\n", resultingString)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
