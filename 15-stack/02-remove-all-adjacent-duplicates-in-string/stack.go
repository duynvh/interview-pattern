package main

import "strconv"

type Stack struct {
	s    []map[rune]int
	size int
}

func (this Stack) Empty() bool {
	return this.size == 0
}

func (this *Stack) Push(data rune, c int) {
	this.s = append(this.s, map[rune]int{data: c})
	this.size++
}

func (this *Stack) Pop() {
	if this.Empty() {
		return
	}

	this.s = this.s[:this.size-1]
	this.size--
}

func (this *Stack) Top() map[rune]int {
	if this.Empty() {
		return make(map[rune]int)
	}

	return this.s[this.size-1]
}

// Size() function will return the size of the array
func (this *Stack) Size() int {
	return this.size
}

/* String() function will convert the stack to string */
func (this *Stack) String() string {
	res := "["
	for i, maps := range this.s {
		for key, data := range maps {
			// Note: This loop will only be executed one time
			// with O(1) time complexity
			res += "['" + string(key) + "', " + strconv.Itoa(data) + "]"
		}
		if i < len(this.s)-1 {
			res += ", "
		}
	}
	res += "]"
	return res
}
