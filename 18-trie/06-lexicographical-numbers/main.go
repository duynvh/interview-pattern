package main

import (
	"fmt"
	"strconv"
)

func Print(node *TrieNode, arr *[]int) {
	if node.completeString {
		*arr = append(*arr, node.data)
	}

	for i := 0; i < 10; i++ {
		str := strconv.Itoa(i)
		if node.children[rune(str[0])] != nil {
			Print(node.children[rune(str[0])], arr)
		}
	}
}

func lexicalOrder(n int) []int {
	trie := NewTrie()

	for i := 1; i <= n; i++ {
		trie.Insert(i)
	}

	result := make([]int, 0)
	Print(trie.root, &result)
	return result
}

func main() {
	result := lexicalOrder(13)
	fmt.Println(result)
}
