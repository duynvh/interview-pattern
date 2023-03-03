package main

// The following class is used for creating nested lists.
// You should not implement this class or speculate about its implementation.
type NestedIntegers struct {
	list    []*NestedIntegers
	integer interface{}
}

// New constructor initializes an empty nested list.
func NewNestedIntegers() *NestedIntegers {
	n := new(NestedIntegers)
	n.list = make([]*NestedIntegers, 0)
	n.integer = nil
	return n
}

// Init constructor initializes a single integer.
func InitNestedIntegers(value interface{}) *NestedIntegers {
	n := new(NestedIntegers)
	n.integer = value
	return n
}

// IsInteger return true if this NestedIntegers holds a single integer,
// rather than a nested list.
func (this *NestedIntegers) IsInteger() bool {
	if this.integer != nil {
		return true
	}
	return false
}

// GetInteger return the single integer that this NestedIntegers holds, if
// it holds a single integer
// Return null if this NestedIntegers holds a nested list
func (this *NestedIntegers) GetInteger() interface{} {
	return this.integer
}

// SetInteger set this NestedIntegers to hold a single integer.
func (this *NestedIntegers) SetInteger(value interface{}) {
	this.list = nil
	this.integer = value
}

// Add set this NestedIntegers to hold a nested list and
// adds a nested integer to it.
func (this *NestedIntegers) Add(ni *NestedIntegers) {
	if this.integer != nil {
		this.list = make([]*NestedIntegers, 0)
		this.list = append(this.list, InitNestedIntegers(this.integer))
		this.integer = nil
	}
	this.list = append(this.list, ni)
}

// GetList return the nested list that this NestedIntegers holds, 
// if it holds a nested list
// Return null if this NestedIntegers holds a single integer
func (this *NestedIntegers) GetList() []*NestedIntegers {
	return this.list
}