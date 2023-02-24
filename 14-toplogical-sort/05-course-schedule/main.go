package main

import (
	"fmt"
	"strings"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	sortedOrder := make([]int, 0)
    if numCourses <= 0 {
        return true
    }

    inDegree := make(map[int]int, 0)
    graph := make(map[int][]int, 0)

    for i := 0; i < numCourses; i++ {
        inDegree[i] = 0
        graph[i] = make([]int, 0)
    }

    // b. Build the graph
    for _, edge := range prerequisites {
        parent, child := edge[1], edge[0]
        graph[parent] = append(graph[parent], child)
        inDegree[child] += 1
    }

    sources := new(Queue)
    for index, _ := range inDegree {
        if inDegree[index] == 0 {
            sources.Enqueue(index)
        }
    }

    for !sources.IsEmpty() {
        course := sources.Dequeue().(int)
        sortedOrder = append(sortedOrder, course)

        for _, child := range graph[course] {
            inDegree[child] -= 1
            if inDegree[child] == 0 {
                sources.Enqueue(child)
            }
        }
    }

    if len(sortedOrder) != numCourses {
        return false
    }

    return true
}

// Driver code
func main() {
	prereq := [][][]int{
		{
			{1, 0},
			{2, 1},
		},
		{
			{1, 0},
			{0, 1},
		},
		{
			{1, 0},
			{2, 1},
			{3, 2},
			{4, 3},
		},
		{
			{1, 0},
			{2, 1},
			{3, 2},
			{4, 3},
			{0, 4},
		},
		{
			{2, 0},
			{2, 1},
			{3, 2},
			{4, 2},
			{3, 1},
		},
	}
	courses := []int{3, 2, 10, 5, 5}

	for index, value := range courses {
		fmt.Printf("%d.\tNumber of courses: %d\n", index+1, value)
		fmt.Printf("\tNumber of pre-requisites: %s\n", strings.Replace(fmt.Sprint(prereq[index]), " ", ", ", -1))
		fmt.Printf("\tOutput: %t\n", canFinish(value, prereq[index]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
