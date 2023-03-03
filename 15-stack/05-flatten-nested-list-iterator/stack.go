package main

import "fmt"

type Stack struct {
	s    []*NestedIntegers
	size int
}

func (this *Stack) Push(v *NestedIntegers) {
	this.s = append(this.s, v)
	this.size++
}

func (this *Stack) Pop() {
	if this.Empty() {
		fmt.Println("StackEmptyException")
		return
	}

    this.s = this.s[:this.size-1]
    this.size--
}

func (this *Stack) Top() *NestedIntegers {
    if this.Empty() {
        fmt.Println("StackEmptyException")
        return nil
    }

    return this.s[this.size-1]
}

/*
	Empty() function will return true if size of the array is

equal to zero or false in all other cases
*/
func (this *Stack) Empty() bool {
	return this.size == 0
}

// Size() function will return the size of the array
func (this *Stack) Size() int {
	return this.size
}
