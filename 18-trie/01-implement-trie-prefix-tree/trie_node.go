package main

type TrieNode struct {
    children map[rune]*TrieNode
    isWord bool
}

func NewTrieNode() *TrieNode {
    node := new(TrieNode)
    node.isWord = false
    node.children = make(map[rune]*TrieNode, 0)
    return node
}