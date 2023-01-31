package main

import "fmt"

func jumpGame(nums []int) int {
	jumps, curEnd, curFarthest := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		curFarthest = max(curFarthest, i+nums[i])
		if i == curEnd {
			jumps++
			curEnd = curFarthest
		}
	}

	return jumps
}

func max(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1
}

func main() {
	//
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jumpGame(nums))
}
