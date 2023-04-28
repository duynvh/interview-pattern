package main

import (
	"fmt"
	"strings"
)

func redundantConnections(edges [][]int) []int {
	graph := new(UnionFind)
	graph.UnionFindInit(len(edges))

	for _, v := range edges {
		if !graph.Union(v[0], v[1]) {
			return []int{v[0], v[1]}
		}
	}

	return []int{-99, -99}
}

// Driver code
func main() {
	edges := [][][]int{
		{
			{1, 2},
			{1, 3},
			{2, 3},
		},
		{
			{1, 2},
			{2, 3},
			{1, 3},
		},
		{
			{1, 2},
			{2, 3},
			{3, 4},
			{1, 4},
			{1, 5},
		},
		{
			{1, 2},
			{1, 3},
			{1, 4},
			{3, 4},
			{2, 4},
		},
		{
			{1, 2},
			{1, 3},
			{1, 4},
			{1, 5},
			{2, 3},
			{2, 4},
			{2, 5},
		},
	}

	for i, v := range edges {
		fmt.Printf("%d.\tEdges: %s\n", i+1, strings.Replace(fmt.Sprint(v), " ", ", ", -1))
		v := redundantConnections(v)
		fmt.Printf("\n\tThe redundant connection in the graph is : [%d, %d]\n", v[0], v[1])
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
