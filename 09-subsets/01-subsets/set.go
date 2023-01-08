package main

import "strconv"

type Set struct {
	hashMap map[int]bool
}

func NewSet() *Set {
	s := new(Set)
	s.hashMap = make(map[int]bool)
	return s
}

func (s *Set) Add(value int) {
	s.hashMap[value] = true
}

// Delete will delete the value from the set.
func (s *Set) Delete(value int) {
	delete(s.hashMap, value)
}

// Exists will check if the value exists in the set or not.
func (s *Set) Exists(value int) bool {
	_, ok := s.hashMap[value]
	return ok
}

/*
	DictToList will convert the list of dictionary to list of list.

DictToList is used as an helper function in a main function for
printing purposes.
*/
func DictToList(sets []Set) [][]int {
	result := make([][]int, len(sets))

	for i := 0; i < len(sets); i++ {
		for k := range sets[i].hashMap {
			result[i] = append(result[i], k)
		}
	}

	return result
}

/*
	DoubleArrayToString is used to convert a 2D integer array to string.

DoubleArrayToString is used as an helper function in a main function for
printing purposes.
*/
func DoubleArrayToString(numsB [][]int) string {
	out := "["
	for i, nums := range numsB {
		out += "["
		for j, num := range nums {
			out += strconv.Itoa(num)
			if j < len(nums)-1 {
				out += ", "
			}
		}
		out += "]"
		if i < len(numsB)-1 {
			out += ", "
		}
	}
	out += "]"
	return out
}

/*
	ArrayToString is used to convert an integer array to string.

ArrayToString is used as an helper function in a main function for
printing purposes.
*/
func ArrayToString(nums []int) string {
	out := "["
	for j, num := range nums {
		out += strconv.Itoa(num)
		if j < len(nums)-1 {
			out += ", "
		}
	}
	out += "]"
	return out
}

// SetToList converts a set to list in order to print it on screen
func SetToList(s Set) []int {
	result := make([]int, 0)

	for k := range s.hashMap {
		result = append(result, k)
	}

	return result
}
