package main

import (
	"fmt"
	"strings"
)

func helper(query string, dict []string, result map[string][]string) []string {
	if query == "" {
		return make([]string, 0)
	}

	for index, _ := range result {
		if index == query {
			return result[query]
		}
	}

	res := make([]string, 0)

	for _, value := range dict {
		if !strings.HasPrefix(query, value) {
			continue
		}

		if len(value) == len(query) {
			res = append(res, value)
		} else {
			resultsOfTheRest := helper(query[len(value):], dict, result)
			for _, v := range resultsOfTheRest {
				v = value + " " + v
				res = append(res, v)
			}
		}
	}

	result[query] = res
	return res
}

// wordBreak is the challenge function
func wordBreak(s string, wordDict []string) []string {
	result := make(map[string][]string)

	// Calling the helper function
	return helper(s, wordDict, result)
}

// Driver code
func main() {
	s := []string{
		// "vegancookbook",
		"catsanddog",
		// "highwaycrash",
		// "pineapplepenapple",
		// "screamicecream",
	}
	subs := []string{
		"cat", "and", "cats", "sand", "dog",
		// "an", "book", "car", "cat", "cook", "cookbook", "crash",
		// "cream", "high", "highway", "i", "ice", "icecream", "low",
		// "scream", "veg", "vegan", "way", "hello", "from", "what",
		// "cats", "and", "dog", "sand", "pineapple", "pine", "apple",
		// "pen", "applepen",
	}

	fmt.Printf("The list of the words we use to break down the strings are:\n")
	fmt.Printf("['%s']\n\n", strings.Join(subs, "', '"))

	for index, value := range s {
		fmt.Printf("%d.\tThe possible strings from the string %s are the following combinations:\n", index+1, s[index])
		fmt.Printf("\t['" + strings.Join(wordBreak(value, subs), `', '`) + "']\n")
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
