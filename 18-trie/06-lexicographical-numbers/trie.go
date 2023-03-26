package main

import (
	"strconv"
)

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	trie := new(Trie)
	trie.root = NewTrieNode()
	return trie
}

func (t *Trie) Insert(key int) {
	node := t.root
	str := strconv.Itoa(key)
	for _, c := range str {
		if _, ok := node.children[c]; !ok {
			node.children[c] = NewTrieNode()
		}

		node = node.children[c]
	}

	node.completeString = true
	node.data = key
}
