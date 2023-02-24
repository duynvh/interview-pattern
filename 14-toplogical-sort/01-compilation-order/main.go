package main

import (
	"fmt"
	"strings"
)

func findOrder(dependencies [][]rune) []rune {
	sortedOrder := make([]rune, 0)

	graph := make(map[rune][]rune)
	inDegree := make(map[rune]int)

	for _, x := range dependencies {
		parent, child := x[1], x[0]
		graph[parent], graph[child] = make([]rune, 0), make([]rune, 0)
		inDegree[parent], inDegree[child] = 0, 0
	}

	if len(graph) <= 0 {
		return sortedOrder
	}

	for _, dependency := range dependencies {
		parent, child := dependency[1], dependency[0]
		graph[parent] = append(graph[parent], child)
		inDegree[child] += 1
	}

	sources := new(Queue)
	for key, _ := range inDegree {
		if inDegree[key] == 0 {
			sources.Enqueue(key)
		}
	}

	for !sources.IsEmpty() {
		vertex := sources.Dequeue().(rune)
		sortedOrder = append(sortedOrder, vertex)
		for _, child := range graph[vertex] {
			inDegree[child] -= 1
			if inDegree[child] == 0 {
				sources.Enqueue(child)
			}
		}
	}

	if len(sortedOrder) != len(graph) {
		return make([]rune, 0)
	}
	return sortedOrder
}

// Driver code
func main() {
	dependencies := [][][]rune{
		{
			{'B', 'A'},
			{'C', 'A'},
			{'D', 'C'},
			{'E', 'D'},
			{'E', 'B'},
		},
		// {
		// 	{'B', 'A'},
		// 	{'C', 'A'},
		// 	{'D', 'B'},
		// 	{'E', 'B'},
		// 	{'E', 'D'},
		// 	{'E', 'C'},
		// 	{'F', 'D'},
		// 	{'F', 'E'},
		// 	{'F', 'C'},
		// },
		// {
		// 	{'A', 'B'},
		// 	{'B', 'A'},
		// },
		// {
		// 	{'B', 'C'},
		// 	{'C', 'A'},
		// 	{'A', 'F'},
		// },
		// {
		// 	{'C', 'C'},
		// },
	}

	for i, prerequisite := range dependencies {
		fmt.Printf("%d.\tdependencies: %s\n", i+1, RuneDoubleArrayToString(prerequisite))
		result := findOrder(prerequisite)
		fmt.Printf("\n\tCompilation Order: %s\n", RuneArrayToString(result))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
