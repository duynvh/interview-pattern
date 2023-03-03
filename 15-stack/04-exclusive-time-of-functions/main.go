package main

import (
	"fmt"
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	stack := new(Stack)
	result := make([]int, n)

	for _, content := range logs {
		log := LogInitWithString(content)

		if log.isStart {
			stack.Push(log)
		} else {
			top := stack.Top().(*Log)
			stack.Pop()

			result[top.id] += (log.time - top.time) + 1

			if stack.Size() > 0 {
				result[stack.Top().(*Log).id] -= (log.time - top.time + 1)
			}
		}
	}

	return result
}

/*
	arrayToString is used to convert an integer array to string.

ArrayToString is used as an helper function in a main function for
printing purposes.
*/
func arrayToString(nums []int) string {
	if len(nums) == 0 {
		return "[]"
	}
	res := "["
	for _, num := range nums {
		res += strconv.Itoa(num) + ", "
	}
	res += "]"
	return res[:len(res)-3] + "]"
}

func main() {

	events := [][]string{
        {"0:start:0", "1:start:2", "1:end:5", "0:end:6"},
        // {"0:start:0", "1:start:2", "1:end:3", "2:start:4", "2:end:7", "0:end:8"},
		// {"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"},
		// {"0:start:0", "1:start:5", "1:end:6", "0:end:7"},
		// {"0:start:0", "1:start:5", "2:start:8", "3:start:12", "4:start:15", "5:start:19", "5:end:22", "4:end:24", "3:end:27", "2:end:32", "1:end:35", "0:end:36"},
		// {"0:start:0", "1:start:3", "1:end:6", "0:end:10"}
    }
    
	n := []int{
        2, 
        // 3, 
        // 2, 
        // 2, 
        // 6, 
        // 2,
    }

	for i, _ := range n {
		fmt.Printf("%d.\tn = %d\n", i+1, n[i])
		time := exclusiveTime(n[i], events[i])
		output := "[" + "\"" + strings.Join(events[i], `", "`) + "\"" + "]"
		fmt.Printf("\tevents = %s\n", output)
		fmt.Printf("\tOutput: %s\n", arrayToString(time))
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}
}
