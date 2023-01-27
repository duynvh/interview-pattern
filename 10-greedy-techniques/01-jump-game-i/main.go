package main

func jumpGame(nums []int) bool {
	targetNumIdx := len(nums) - 1
	for i := len(nums) - 2; i > -1; i-- {
		if targetNumIdx <= i+nums[i] {
			targetNumIdx = i
		}
	}

	return targetNumIdx == 0
}
