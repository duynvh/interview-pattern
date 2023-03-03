package main

import "strconv"

type Stack struct {
	s    []int
	size int
}

func (this *Stack) Size() int {
	return this.size
}

func (this *Stack) Empty() bool {
	return this.size == 0
}

func (this *Stack) Push(v int) {
	this.s = append(this.s, v)
	this.size++
}

func (this *Stack) Pop() {
	if this.Empty() {
		return
	}

	this.s = this.s[:this.size-1]
	this.size--
}

func (this *Stack) Top() int {
	if this.Empty() {
		return 0
	}

	return this.s[this.size-1]
}

// String converts stack to string
// String function is used for printing purposes
func (this *Stack) String() string {
	res := "["
	for i, num := range this.s {
		res += strconv.Itoa(num)
		if i < len(this.s)-1 {
			res += ", "
		}
	}
	res += "]"
	return res
}
