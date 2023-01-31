package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) [][]rune {
	output, result := make([]rune, 0), make([][]rune, 0)
	backtrack(n, 0, 0, &output, &result)
	return result
}

func backtrack(n int, openBracesCount int, closeBracesCount int, output *[]rune, result *[][]rune) {
	if openBracesCount >= n && closeBracesCount >= n {
		outputStr := make([]rune, len(*output))
		copy(outputStr, *output)
		*result = append(*result, outputStr)
	}

	// Case where we can still add opening braces
	if openBracesCount < n {
		*output = append(*output, '(')
		backtrack(n, openBracesCount+1, closeBracesCount, output, result)
		*output = (*output)[:len(*output)-1]
	}

	// Case where we add closing braces if the current count of closing braces is less than the count of opeing braces
	if closeBracesCount < openBracesCount {
		*output = append(*output, ')')
		backtrack(n, openBracesCount, closeBracesCount+1, output, result)
		*output = (*output)[:len(*output)-1]
	}
}

/*
	PrintResult function is used to print the result vector in an easy

to view manner from main, not part of the actual function
*/
func printResult(res [][]rune) {
	for _, value := range res {
		fmt.Printf("\t\t%s\n", string(value))
	}
}

func main() {
	inputList := []int{
		// 1,
		// 2,
		3,
		// 4,
		// 5,
	}

	for index, value := range inputList {
		fmt.Printf("%d.\tn = %d\n", index+1, value)
		fmt.Print("\tAll combinations of valid balanced parentheses:\n")
		result := generateParenthesis(value)
		printResult(result)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
