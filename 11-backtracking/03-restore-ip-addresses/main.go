package main

import (
	"fmt"
	"strconv"
	"strings"
)

// valid is used as helper function
func valid(segment string) bool {
	segmentLength := len(segment) // storing the length of each segment
	if segmentLength > 3 {        // each segment's length should be less than 3
		return false
	}
	// Check if the current segment is valid
	// for either one of following conditions:
	// 1. Check if the current segment is less or equal to 255.
	// 2. Check if the length of segment is 1. The first character of
	//    segment can be `0` only if the length of segment is 1.
	if segment[0] != '0' {
		val, _ := strconv.Atoi(segment)
		return val <= 255
	} else {
		return len(segment) == 1
	}
}

// updateSegment function will append the current list of segments to
// the list of result.
func updateSegment(s string, currPos int, segments *[]string, result *[]string) {
	segment := s[currPos+1 : len(s)]

	if valid(segment) { // if the segment is acceptable
		// add it to the list of segments
		*segments = append(*segments, segment)
		*result = append(*result, strings.Join(*segments, "."))
		// remove the top segment
		*segments = (*segments)[:len(*segments)-1]
	}
}

// min function returns the minimum of the integers provided
func min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

// backtrack is used as helper function
func backtrack(s string, prevPos int, dots int, segments *[]string, result *[]string) {
	// prevPos : the position of the previously placed dot
	// dots : number of dots to place
	size := len(s)

	// The current dot currPos could be placed in
	// a range from prevPos + 1 to prevPos + 4.
	// The dot couldn't be placed after the last character in the string.
	for currPos := prevPos + 1; currPos < min(size-1, prevPos+4); currPos++ {
		segment := s[prevPos+1 : currPos+1]
		if valid(segment) {
			*segments = append(*segments, segment)

			// if all 3 dots are placed add the solution to result
			if dots-1 == 0 {
				updateSegment(s, currPos, segments, result)
			} else {
				// continue to place dots
				backtrack(s, currPos, dots-1, segments, result)
			}
			// remove the last placed dot
			*segments = (*segments)[:len(*segments)-1]
		}
	}
}

// restoreIpAddresses is the solution function
func restoreIpAddresses(s string) []string {
	// creating empty lists for storing valid IP addresses,
	// and each segment of IP
	result, segments := make([]string, 0), make([]string, 0)
	backtrack(s, -1, 3, &segments, &result)
	return result
}

/*
	stringArrayToString() is used to convert a string array to

string. It is used for printing purposes in driver code.
*/
func stringArrayToString(stringArray []string) string {
	out := strings.Join(stringArray, "\", \"")
	out = "[\"" + out[0:1] + out[1:len(out)] + "\"]"
	return out
}

// driver code
func main() {
	// input streams of IP addresses
	ipAddresses := []string{
		// "0000",
		"25525511135",
		// "12121212",
		// "113242124",
		// "199219239",
		// "121212",
		// "25525511335",
	}

	// loop to execute till the lenght of input IP addresses
	for i, ipAddress := range ipAddresses {
		fmt.Printf("%d.\t Input addresses: \"%s\"\n", i+1, ipAddress)
		res := restoreIpAddresses(ipAddress)
		fmt.Printf("\tPossible valid IP Addresses are: %s\n", stringArrayToString(res))
		fmt.Printf("\n\n%s\n", strings.Repeat("-", 100))
	}
}
