package main

import "sort"

func makeSquare(matchsticks []int) bool {
	sum, target := 0, 0
	for i := 0; i < len(matchsticks); i++ {
		sum += matchsticks[i]
	}

	target = sum / 4
	if sum != target*4 {
		return false
	}

	// optimize performance
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})

    sides := [4]int{}
    var dfs func(i int) bool
    dfs = func(i int) bool {
        if i == len(matchsticks) {
            return true
        }

        for j := 0; j < 4; j++ {
            if sides[j]+matchsticks[i] > target {
                continue
            }

            sides[j] += matchsticks[i]
            if dfs(i + 1) {
                return true
            }
            sides[j] -= matchsticks[i]
        }

        return false
    }

    return dfs(0)
}
