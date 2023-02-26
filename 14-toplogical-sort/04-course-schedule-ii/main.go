package main

import (
	"container/list"
	"fmt"
	"strings"
)

func findOrder(n int, prerequisites [][]int) []int {
	var sortedOrder []int

	if n <= 0 {
		return sortedOrder
	}

	inDegree := make(map[int]int)

	// Adjacency list graph
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		inDegree[i] = 0
		graph[i] = make([]int, 0)
	}

	// b. Build the graph
	for i := 0; i < len(prerequisites); i++ {
		parent, child := prerequisites[i][1], prerequisites[i][0]
		graph[parent] = append(graph[parent], child)
		inDegree[child] = inDegree[child] + 1
	}

	// c. Find all source i.e, all n with 0 in-degrees
	sources := list.New()
	for key, val := range inDegree {
		if val == 0 {
			sources.PushBack(key)
		}
	}

	for sources.Len() != 0 {
		front := sources.Front()
		vertex := front.Value.(int)
		sources.Remove(front)
		sortedOrder = append(sortedOrder, vertex)

		children := graph[vertex]
		for _, child := range children {
			inDegree[child] = inDegree[child] - 1
			if inDegree[child] == 0 {
				sources.PushBack(child)
			}
		}
	}

	if len(sortedOrder) != n {
		sortedOrder = []int{}
	}

	return sortedOrder
}

// Driver Code
func main() {
	n := []int{4, 5, 0, 4, 7}
	prerequisites := [][][]int{
		{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
		{{1, 0}, {2, 0}, {3, 1}, {4, 3}}, {{1, 0}},
		{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
		{{1, 0}, {0, 3}, {0, 2}, {3, 2}, {2, 5}, {4, 5}, {5, 6}, {2, 4}},
	}

	for i, prerequisite := range prerequisites {
		fmt.Printf("%d.\tPrerequisites: %s\n", i+1, strings.Replace(fmt.Sprint(prerequisite), " ", ", ", -1))
		fmt.Printf("\tTotal number of course, n: %d\n", n[i])
		fmt.Printf("\tValid courses order: %s\n", strings.Replace(fmt.Sprint(findOrder(n[i], prerequisite)), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
