package main

import "sort"

func twoCityScheduling(costs [][]int) int {
    // Array that we will use to divide the group in two equal parts
    difference := make([][]int, 0)

    // Initilizing a new variable to calculate the minimum cost required to send exactly n people to both cities
    answer := 0

    // Initiliazing a loop, to calculate the different to travel to City A or B
    for _, cost := range costs {
        difference = append(difference, []int{cost[0] - cost[1], cost[0], cost[1]})
    }

    // We sort the array based on the differences we calculated above
    sort.Slice(difference[:], func(i, j int) bool {
        for x := range difference[i] {
            if difference[i][x] == difference[j][x] {
                continue
            }
            return difference[i][x] < difference[j][x]
        }
        return false
    })

    lengthDiff := len(difference)
    // Initiliazing a loop, to add the relevent costs to our answer variable
    for i := range difference {
        if i < (lengthDiff/2) {
            // We will send this half to City A
            answer = answer + difference[i][1]
        } else {
            // We will send this half to City B
            answer = answer + difference[i][2]
        }
    }

    return answer
}