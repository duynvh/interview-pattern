package main

type TrieNode struct {
    children map[rune]*TrieNode
    isString bool
}

func NewTrieNode() *TrieNode {
    node := new(TrieNode)
    node.children = make(map[rune]*TrieNode, 0)
    node.isString = false
    return node
}