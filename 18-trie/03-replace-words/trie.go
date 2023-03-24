package main

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    trie := new(Trie)
    trie.root = NewTrieNode()
    return trie
}

func (t *Trie) Insert(word string) {
    node := t.root

    for _, c := range word {
        if _, ok := node.children[c]; !ok {
            node.children[c] = NewTrieNode()
        }

        node = node.children[c]
    }

    node.endOfWord = true
}

func (t *Trie) Replace(word string) string {
    curr := t.root
    for i, c := range word {
        if _, ok := curr.children[c]; !ok {
            return word
        }

        curr = curr.children[c]

        if curr.endOfWord == true {
            return word[:i+1]
        }
    }

    return word
}