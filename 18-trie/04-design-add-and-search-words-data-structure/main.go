package main

import (
	"fmt"
	"strings"
)

type WordDictionary struct {
	root    *TrieNode
	canFind bool
}

func NewWordDictionary() *WordDictionary {
	dict := new(WordDictionary)
	dict.root = NewTrieNode()
	dict.canFind = false
	return dict
}

func (w *WordDictionary) AddWord(word string) {
	n := len(word)
	curNode := w.root
	for i, val := range word {
		index := int(val) - int('a')
		if curNode.nodes[index] == nil {
			curNode.nodes[index] = NewTrieNode()
		}
		curNode = curNode.nodes[index]

		if i == n-1 {
			if curNode.complete {
				return
			}

			curNode.complete = true
		}
	}
}

func (w *WordDictionary) SearchWord(word string) bool {
	w.canFind = false
	w.DepthFirstSearch(w.root, word, 0)
	return w.canFind
}

func (w *WordDictionary) DepthFirstSearch(root *TrieNode, word string, i int) {
	if w.canFind {
		return
	}

	if root == nil {
		return
	}

	if len(word) == i {
		if root.complete {
			w.canFind = true
		}

		return
	}

	if word[i] == '.' {
		for j := int('a'); j < int('z'); j++ {
			w.DepthFirstSearch(root.nodes[j-int('a')], word, i+1)
		}
	} else {
		index := int(word[i]) - int('a')
		w.DepthFirstSearch(root.nodes[index], word, i+1)
	}
}

func (w *WordDictionary) GetWords() []string {
	wordsList := make([]string, 0)
	if w.root == nil {
		return make([]string, 0)
	}

	return w.Dfs(w.root, "", wordsList)
}

func (w *WordDictionary) Dfs(node *TrieNode, word string, wordList []string) []string {
	if node == nil {
		return wordList
	}

	if node.complete {
		wordList = append(wordList, word)
	}

	for i := int('a'); i < int('z'); i++ {
		n := word + string(rune(i))
		wordList = w.Dfs(node.nodes[i-int('a')], n, wordList)
	}

	return wordList
}

/*
	stringArrayToString() is used to convert a string array to

string. It is used for printing purposes in driver code.
*/
func stringArrayToString(str []string) string {
	res := "["
	for j, s := range str {
		res += "\"" + s + "\""
		if j < len(str)-1 {
			res += ", "
		}
	}
	res += "]"
	return res
}

func main() {
	obj := NewWordDictionary()
	// Adding words
	words := []string{"add", "hello", "answer", "add"}
	i := 1
	for _, w := range words {
		fmt.Printf("%d.\tAdding word: \"%s\"\n", i, w)
		obj.AddWord(w)
		fmt.Printf("%s\n", strings.Repeat("-", 93))
		i++
	}

	// Searching words
	wordSearch := []string{"hello", "xyz", ".lo", "...lo"}
	for _, w := range wordSearch {
		fmt.Printf("%d.\tSearching word: \"%s\"\n", i, w)
		fmt.Printf("\t%v\n", obj.SearchWord(w))
		fmt.Printf("%s\n", strings.Repeat("-", 93))
		i += 1
	}
	// Getting all words
	fmt.Printf("%d.\tGetting all words: %s\n", i, stringArrayToString(obj.GetWords()))
	fmt.Printf("%s\n", strings.Repeat("-", 93))
}
