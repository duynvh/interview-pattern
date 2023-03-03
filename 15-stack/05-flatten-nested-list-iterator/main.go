package main

import (
	"fmt"
	"strings"
)

type NestedIterator struct {
	nestedListStack Stack
}

func Constructor(nestedList []*NestedIntegers) *NestedIterator {
	n := new(NestedIterator)

	for i := len(nestedList) - 1; i >= 0; i-- {
		n.nestedListStack.Push(nestedList[i])
	}

	return n
}

func (this *NestedIterator) HasNext() bool {
	for !this.nestedListStack.Empty() {
		top := this.nestedListStack.Top()

		if top.IsInteger() {
			return true
		}

		topList := top.GetList()
		this.nestedListStack.Pop()
		for i := len(topList) - 1; i >= 0; i-- {
			this.nestedListStack.Push(topList[i])
		}
	}

	return false
}

func (this *NestedIterator) Next() int {
	if this.HasNext() {
		top := this.nestedListStack.Top()
		this.nestedListStack.Pop()
		return top.GetInteger().(int)
	}

	return -1
}

// main is used for driver code
func main() {
	inputs := []string{"[1, [2, 3], 4]",
		"[3, [2, 3, 4], 4, [2, 3]]",
		"[[2, 3], 3, [2, 3], 4, [2, 3, 4, 5]]",
		"[1, [3, [4, [5, 6], 7], 8], 9]",
		"[[2, 3, [2, 3]]]"}
	itr := make([]*NestedIterator, 0)

	// Test Case #1
	nestedList := make([]*NestedIntegers, 0)
	l1 := NewNestedIntegers()
	nestedList = append(nestedList, InitNestedIntegers(1))
	l1.Add(InitNestedIntegers(2))
	l1.Add(InitNestedIntegers(3))
	nestedList = append(nestedList, l1)
	nestedList = append(nestedList, InitNestedIntegers(4))
	itr = append(itr, Constructor(nestedList))

	// Test Case #2
	nestedList = make([]*NestedIntegers, 0)
	l2 := NewNestedIntegers()
	nestedList = append(nestedList, InitNestedIntegers(3))
	l2.Add(InitNestedIntegers(2))
	l2.Add(InitNestedIntegers(3))
	l2.Add(InitNestedIntegers(4))
	nestedList = append(nestedList, l2)
	nestedList = append(nestedList, InitNestedIntegers(4))
	nestedList = append(nestedList, l1)
	itr = append(itr, Constructor(nestedList))

	// Test Case #3
	nestedList = make([]*NestedIntegers, 0)
	l3 := NewNestedIntegers()
	nestedList = append(nestedList, l1)
	nestedList = append(nestedList, InitNestedIntegers(3))
	l3.Add(InitNestedIntegers(2))
	l3.Add(InitNestedIntegers(3))
	l3.Add(InitNestedIntegers(4))
	l3.Add(InitNestedIntegers(5))
	nestedList = append(nestedList, l1)
	nestedList = append(nestedList, InitNestedIntegers(4))
	nestedList = append(nestedList, l3)
	itr = append(itr, Constructor(nestedList))

	// Test case #4
	nestedList = make([]*NestedIntegers, 0)
	nestedList = append(nestedList, InitNestedIntegers(1))
	l1 = NewNestedIntegers()
	l1.Add(InitNestedIntegers(5))
	l1.Add(InitNestedIntegers(6))
	l2 = NewNestedIntegers()
	l2.Add(InitNestedIntegers(4))
	l2.Add(l1)
	l2.Add(InitNestedIntegers(7))
	l3 = NewNestedIntegers()
	l3.Add(InitNestedIntegers(3))
	l3.Add(l2)
	l3.Add(InitNestedIntegers(8))
	nestedList = append(nestedList, l3)
	nestedList = append(nestedList, InitNestedIntegers(9))
	itr = append(itr, Constructor(nestedList))

	// TEST CASE 5: [[2, 3, [2, 3]]]
	nestedList = make([]*NestedIntegers, 0)
	l2 = NewNestedIntegers()
	l2.Add(InitNestedIntegers(2))
	l2.Add(InitNestedIntegers(3))
	l3 = NewNestedIntegers()
	l3.Add(InitNestedIntegers(2))
	l3.Add(InitNestedIntegers(3))
	l3.Add(l2)
	nestedList = append(nestedList, l3)
	itr = append(itr, Constructor(nestedList))

	for i := range itr {
		fmt.Printf("%d.\tOriginal structure: %s\n", i+1, inputs[i])
		fmt.Printf("\n\tOutput:\n")
		for itr[i].HasNext() {
			fmt.Printf("\titr.next(): %d\n", itr[i].Next())
		}
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}
}
