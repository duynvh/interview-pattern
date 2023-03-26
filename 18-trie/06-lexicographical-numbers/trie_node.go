package main

type TrieNode struct {
    children map[rune]*TrieNode
    completeString bool
    data int
}

func NewTrieNode() *TrieNode {
    node := new(TrieNode)
    node.children = make(map[rune]*TrieNode, 0)
    node.completeString = false
    node.data = -1
    return node
}