package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	root *TrieNode
}

func constructor() *Trie {
	t := new(Trie)
	t.root = NewTrieNode()
	return t
}

// Insert function insert string in trie
func (t *Trie) Insert(word string) {
	node := t.root
	// adding string characters in the tree
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			node.children[c] = NewTrieNode()
		}
		node = node.children[c] // update the node as we move to the next character
	}
	node.isWord = true // set isWord as true when all the string characters have been added
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return false
		}

		node = node.children[c]
	}

	return node.isWord
}

func (t *Trie) SearchPrefix(prefix string) bool {
	node := t.root

	for _, c := range prefix {
		if _, ok := node.children[c]; !ok {
			return false
		}

		node = node.children[c]
	}

	return true
}

// main is used for driver code
func main() {
	keys := []string{"the", "a", "there", "answer"}
	trieForKeys := new(Trie)
	trieForKeys.root = NewTrieNode()
	for i, x := range keys {
		fmt.Printf("%d.\tInserting key: \"%s\"", i+1, x)
		trieForKeys.Insert(x)
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}

	search := []string{"a", "answer", "xyz", "an"}
	for _, y := range search {
		fmt.Printf("\tWord found? %v\n", trieForKeys.Search(y))
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}

	searchPrefix := []string{"b", "an"}
	for _, z := range searchPrefix {
		fmt.Printf("\tPrefix Found? %v\n", trieForKeys.SearchPrefix(z))
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}
}
