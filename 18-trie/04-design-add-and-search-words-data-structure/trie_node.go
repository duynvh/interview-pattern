package main

type TrieNode struct {
	nodes    []*TrieNode
	complete bool
}

func NewTrieNode() *TrieNode {
	node := new(TrieNode)
	node.complete = false

	for i := 0; i < 26; i++ {
		node.nodes = append(node.nodes, new(TrieNode))
		node.nodes[i] = nil
	}

	return node
}
