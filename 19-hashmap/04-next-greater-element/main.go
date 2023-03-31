package main

import (
	"fmt"
	"strconv"
	"strings"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hashmap := make(map[int]int) // Initliazing the hashmap
	ans := make([]int, len(nums1))
	for i, n := range nums1 {
		hashmap[n] = i
		ans[i] = -1
	}

	for i := range nums2 {
		if _, ok := hashmap[nums2[i]]; !ok {
			continue
		}

		for j := i + 1; j < len(nums2); j++ {
			if nums2[j] > nums2[i] {
				index, _ := hashmap[nums2[i]]
				ans[index] = nums2[j]

				break
			}
		}
	}

	return ans
}

/*
	arrayToString is used to convert an integer array to string.

arrayToString is used as an helper function in a main function for
printing purposes.
*/
func arrayToString(nums []int) string {
	res := "["
	for i, num := range nums {
		res += strconv.Itoa(num)
		if i < len(nums)-1 {
			res += ", "
		}
	}
	res += "]"
	return res
}

func main() {
	nums1 := [][]int{{2, 4}, {3, 2, 5}, {14, 45, 52}, {1, 3, 2}, {4, 2}, {0}}
	nums2 := [][]int{{1, 2, 3, 4}, {2, 3, 5, 1}, {52, 14, 45, 65}, {1, 3, 2, 4, 5}, {1, 2, 4, 3}, {0}}

	for i := range nums1 {
		fmt.Printf("%d.\tNums 1 = %s\n", i+1, arrayToString(nums1[i]))
		fmt.Printf("\tNums 2 = %s\n", arrayToString(nums2[i]))
		result := nextGreaterElement(nums1[i], nums2[i])
		fmt.Printf("\n\tThe Next Greater Element Array = %s\n", arrayToString(result))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
