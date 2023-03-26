package main

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(key string) {
	node := t.root
	for _, c := range key {
		if _, ok := node.children[c]; !ok {
			node.children[c] = &TrieNode{children: make(map[rune]*TrieNode)}
		}

		node = node.children[c]
	}

	node.isString = true
}

func (t *Trie) Search(key string) bool {
	node := t.root
	for _, c := range key {
		if _, ok := node.children[c]; !ok {
			return false
		}

		node = node.children[c]
	}

	return node.isString
}

func (t *Trie) startsWith(prefix string) bool {
	node := t.root
	for _, c := range prefix {
		if _, ok := node.children[c]; !ok {
			return false
		}

		node = node.children[c]
	}

	return true
}
