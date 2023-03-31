package main

import (
	"fmt"
	"strings"
)

type MyHashMap struct {
	keySpace int
	buckets  []*Bucket
}

func (this *MyHashMap) MyHashMapInit(k int) {
	this.keySpace = k
	this.buckets = make([]*Bucket, k)
	for i := 0; i < k; i++ {
		this.buckets[i] = BucketInit()
	}
}

func (this *MyHashMap) Put(key int, value int) {
	hashKey := key % this.keySpace
	this.buckets[hashKey].Update(key, value)
}

// get function is used to fetch corresponding value of the given key
func (this *MyHashMap) Get(key int) int {
	hashKey := key % this.keySpace
	return this.buckets[hashKey].Get(key)
}

// remove function is used to delete corresponding value of the given key
func (this *MyHashMap) Remove(key int) {
	hashKey := key % this.keySpace
	this.buckets[hashKey].Remove(key)
}

// Driver code
func main() {
	keySpace := 11
	input := new(MyHashMap)
	input.MyHashMapInit(keySpace)
	keysList := []int{5, 11, 12, 15, 22, 10}
	funcs := []string{"Get", "Get", "Put", "Get", "Put", "Get", "Get", "Remove", "Get", "Get", "Remove", "Get"}
	funcKeys := [][]int{
		{5},
		{15},
		{15, 250},
		{15},
		{121, 110},
		{121},
		{10},
		{11},
		{11},
		{13},
		{13},
		{0},
	}

	for i, _ := range funcs {
		if funcs[i] == "Put" {
			fmt.Printf("%d.\tput(%d, %d)\n", i+1, funcKeys[i][0], funcKeys[i][1])
			present := false
			for j, _ := range keysList {
				if funcKeys[i][0] == keysList[j] {
					present = true
				}
			}
			if !present {
				keysList = append(keysList, funcKeys[i][0])
			}
			input.Put(funcKeys[i][0], funcKeys[i][1])
		} else if funcs[i] == "Get" {
			fmt.Printf("%d.\tget(%d)\n", i+1, funcKeys[i][0])
			fmt.Printf("\tValue returned: %d\n", input.Get(funcKeys[i][0]))
		} else if funcs[i] == "Remove" {
			fmt.Printf("%d.\tremove(%d)\n", i+1, funcKeys[i][0])
			input.Remove(funcKeys[i][0])
		}

		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
