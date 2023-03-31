package main

import (
	"fmt"
	"strings"
)

type RequestLogger struct {
	requests map[string]int
	limit    int
}

func (this *RequestLogger) requestLoggerInit(timeLimit int) {
	this.requests = make(map[string]int)
	this.limit = timeLimit
}

func (this *RequestLogger) messageRequestDecision(timestamp int, request string) bool {
	for index, _ := range this.requests {
		if index == request && timestamp-this.requests[request] >= this.limit {
			this.requests[index] = timestamp
			return true
		} else if index == request {
			return false
		}
	}

	this.requests[request] = timestamp
	return true
}

// Driver code
func main() {
	newRequests := new(RequestLogger)
	newRequests.requestLoggerInit(8)
	times := []int{1, 5, 6, 7, 15}
	messages := []string{"good morning", "hello world", "good morning", "good morning", "hello world"}

	for index, value := range messages {
		fmt.Printf("%d.\tTime, Message: {%d, '%s'}\n", index+1, times[index], value)
		fmt.Printf("\tMessage request decision: %t\n", newRequests.messageRequestDecision(times[index], value))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
