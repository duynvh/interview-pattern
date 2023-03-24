package main

type TrieNode struct {
    children map[rune]*TrieNode
    searchWords []string
}

func NewTrieNode() *TrieNode {
    node := new(TrieNode)
    node.searchWords = make([]string, 0)
    node.children = make(map[rune]*TrieNode)
    return node
}