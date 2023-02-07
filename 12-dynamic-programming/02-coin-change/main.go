package main

import (
	"fmt"
	"math"
	"strings"
)

func calculateMinimumCoins(coins []int, rem int, counter []int) int {
	if rem < 0 {
		return -1
	}

	if rem == 0 {
		return 0
	}

	if counter[rem-1] != math.MaxInt32 {
		return counter[rem-1]
	}

	minimum := math.MaxInt32
	for _, value := range coins {
		result := calculateMinimumCoins(coins, rem-value, counter)

		if result >= 0 && result < minimum {
			minimum = 1 + result
		}
	}

	if minimum != math.MaxInt32 {
		counter[rem-1] = minimum
	} else {
		counter[rem-1] = -1
	}

	return counter[rem-1]
}

// func coinChange(coins []int, total int) int {
//     if total < 1 {
//         return 0
//     }

//     counter := make([]int, total)
//     for index, _ := range counter {
//         counter[index] = math.MaxInt32
//     }

//     return calculateMinimumCoins(coins, total, counter)
// }

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = math.MaxInt32
    }

    dp[0] = 0

	for a := 1; a <= amount; a++ {
        for _, c := range coins {
            if a - c < 0 {
                continue
            }

            dp[a] = min(dp[a], dp[a-c] + 1)
        }
    }

	if dp[amount] == math.MaxInt32 {
        return -1
    }

    return dp[amount]
}

func min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	}

	return v2
}

// Driver code
func main() {
	coins := [][]int{
		{1, 3, 4, 5},
		// {1, 4, 6, 9},
		// {6, 7, 8},
		// {1, 2, 3, 4, 5},
		// {14, 15, 18, 20},
	}
	total := []int{
		7,
		// 11,
		// 27,
		// 41,
		// 52,
	}

	for index, value := range coins {
		fmt.Printf("%d.\tThe minimum number of coins required to find %d from %s is: %d\n", index+1, total[index], strings.Replace(fmt.Sprint(value), " ", ", ", -1), coinChange(value, total[index]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
