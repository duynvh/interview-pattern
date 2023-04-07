package main

import (
	"fmt"
	"strings"
)

type FreqStack struct {
	frequency    map[int]int
	group        map[int][]int
	maxFrequency int
}

func Constructor() FreqStack {
	return FreqStack{
		frequency:    make(map[int]int),
		group:        make(map[int][]int),
		maxFrequency: 0,
	}
}

func (this *FreqStack) Push(value int) {
	freq, _ := this.frequency[value]
	freq += 1
	this.frequency[value] = freq

	if freq > this.maxFrequency {
		this.maxFrequency = freq
	}

	this.group[freq] = append(this.group[freq], value)
}

func (this *FreqStack) Pop() int {
	value := 0
	if this.maxFrequency > 0 {
		// Fetch the top of the group[this.maxFrequency] stack
		value = this.group[this.maxFrequency][len(this.group[this.maxFrequency])-1]

		// Pop the top of the group[this.maxFrequency] stack
		this.group[this.maxFrequency] = this.group[this.maxFrequency][:len(this.group[this.maxFrequency]) - 1]
		this.frequency[value]--

		if len(this.group[this.maxFrequency]) == 0 {
			this.maxFrequency--
		}
	} else {
		return -1
	}

	return value
}

func main() {
	obj := Constructor()
	inputs := []int{5, 7, 7, 7, 4, 5, 3}
	fmt.Printf("\tInput Stack: %s\n\n", strings.Replace(fmt.Sprint(inputs), " ", ", ", -1))

	for _, i := range inputs {
		obj.Push(i)
	}

	for i := range inputs {
		fmt.Printf("%d.\tPopping out the most frequent value... \n", i+1)
		fmt.Printf("\tValue removed from stack is: %d\n", obj.Pop())
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}
}
