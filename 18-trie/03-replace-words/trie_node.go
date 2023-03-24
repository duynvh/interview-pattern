package main

type TrieNode struct {
    children map[rune]*TrieNode
    endOfWord bool
}

func NewTrieNode() *TrieNode {
    node := new(TrieNode)
    node.children = make(map[rune]*TrieNode)
    node.endOfWord = false
    return node
}