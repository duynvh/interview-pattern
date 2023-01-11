package main

import "fmt"

func helper(result *[][]int, temp []int, nums []int, index int) {
	if index == len(nums) {
		copyTemp := make([]int, len(temp))
		copy(copyTemp, temp)

		*result = append(*result, copyTemp)
	} else {
		temp = append(temp, nums[index])
		helper(result, temp, nums, index+1)
		temp = temp[:len(temp)-1]
		helper(result, temp, nums, index+1)
	}
}

func subsets(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return result
	}
	helper(&result, []int{}, nums, 0)
	return result
}

func main() {
	inputList := []int{
		1,
		2,
		3,
		// 4,
		// 5,
	}

	result := subsets(inputList)
	for _, value := range result {
		fmt.Println(value)
	}
}
