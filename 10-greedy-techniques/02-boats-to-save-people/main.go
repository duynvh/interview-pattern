package main

import (
	"fmt"
	"sort"
	"strings"
)

func rescueBoats(people []int, limit int) int {
	sort.Ints(people)

	lightest, heaviest := 0, len(people)-1
	answer := 0

	for lightest <= heaviest {
		answer += 1
		if people[lightest]+people[heaviest] <= limit {
			lightest += 1
		}

		heaviest -= 1
	}

	return answer
}

func main() {
	people := [][]int{{1, 2}, {3, 2, 2, 1}, {3, 5, 3, 4}, {
		5, 5, 5, 5}, {1, 2, 3, 4}, {1, 2, 3, 4, 5}, {3, 4, 5}}
	limit := []int{3, 3, 5, 5, 5, 3, 1}
	for i, peopleList := range people {
		fmt.Printf("%d.\tWeights = %s\n", i+1, strings.Replace(fmt.Sprint(peopleList), " ", ", ", -1))
		fmt.Printf("\tWeight Limit = %d\n", limit[i])
		fmt.Printf("\tThe minimum number of boats required to save people are %d\n", rescueBoats(peopleList, limit[i]))
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}
}
