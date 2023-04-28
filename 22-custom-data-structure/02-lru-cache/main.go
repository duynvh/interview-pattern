package main

import (
	"fmt"
	"strconv"
	"strings"
)

type LRUCache struct {
	cacheCapacity int
	cacheMap      map[int]*EduLinkedListNode
	cacheList     *EduLinkedList
}

func LRUCacheInitWithCapacity(capacity int) *LRUCache {
	LRU := &LRUCache{
		cacheCapacity: capacity,
		cacheMap:      make(map[int]*EduLinkedListNode),
		cacheList:     LinkedListInit(),
	}

	return LRU
}

func (l *LRUCache) Get(key int) int {
	foundIter := l.cacheMap[key]
	if foundIter == nil {
		return -1
	}

	listIter := foundIter
	l.cacheList.MoveToHead(foundIter)
	return listIter.second
}

func (l *LRUCache) Set(key int, value int) {
	_, keyExists := l.cacheMap[key]

	if keyExists {
		foundIter := l.cacheMap[key]
		listIter := foundIter

		l.cacheList.MoveToHead(foundIter)
		listIter.pair[1] = value
		return
	}

	if len(l.cacheMap) == l.cacheCapacity {
		keyTemp := l.cacheList.GetTail().pair[0]
		l.cacheList.RemoveTail()
		delete(l.cacheMap, keyTemp)
	}

	l.cacheList.InsertAtHead([2]int{key, value})

	_, haskey := l.cacheMap[key]
	if haskey {
		l.cacheMap[key] = l.cacheList.GetHead()
	} else {
		l.cacheMap[key] = l.cacheList.GetHead()
	}
}

// Print prints the current state of the cache
func (l *LRUCache) Print() {

	fmt.Printf("Cache current size: %d, Cache contents: {", l.cacheList.size)

	node := l.cacheList.GetHead()
	for node != nil {
		fmt.Printf("{%d: %d}", node.pair[0], node.pair[1])
		node = node.next

		if node != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("}")
	fmt.Printf("\n%s\n", strings.Repeat("-", 100))
}

func main() {

	cacheCapacity := 2
	// Creating a cache of size 2
	cache := LRUCacheInitWithCapacity(cacheCapacity)

	fmt.Print("Initial state of cache\n")
	fmt.Printf("Cache capacity: %d\n", cacheCapacity)
	cache.Print()

	keys := [7]int{10, 10, 15, 20, 15, 25, 5}
	values := [7]string{"20", "Get", "25", "40", "Get", "85", "5"}

	for i, key := range keys {
		if values[i] == "Get" {
			fmt.Printf("Getting by Key: %d\n", key)
			fmt.Printf("Cached value returned: %d\n", cache.Get(key))
		} else {
			fmt.Printf("Setting cache: Key: %d, Value: %s\n", key, values[i])
			v, _ := strconv.Atoi(values[i])
			cache.Set(key, v)
		}
		cache.Print()
	}
}
