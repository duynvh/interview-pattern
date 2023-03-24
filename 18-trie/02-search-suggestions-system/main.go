package main

import (
	"fmt"
	"sort"
	"strings"
)

type Trie struct {
	root *TrieNode
}

// doublyStringArrayToString() is used to convert a 2D string
// array to string. It is used for printing purposes.
func doublyStringArrayToString(str [][]string) string {
	out := "["
	for _, v := range str {
		if len(v) == 0 {
			out += "[], "
		} else {
			out += "['" + strings.Join(v, "', '") + "'], "
		}
	}
	out = out[0 : len(out)-2]
	out += "]"
	return out
}

func (this *Trie) Insert(data string) {
	node := this.root

	for _, char := range data {
		if _, ok := node.children[char]; !ok {
			node.children[char] = NewTrieNode()
		}

		node = node.children[char]
		if len(node.searchWords) < 3 {
			node.searchWords = append(node.searchWords, data)
		}
	}
}

func (this *Trie) Search(word string) [][]string {
	result := make([][]string, 0)
	node := this.root

	for i, char := range word {
		if _, ok := node.children[char]; !ok {
			temp := make([][]string, len(word)-i)
			result = append(result, temp...)
			return result
		} else {
			node = node.children[char]
			result = append(result, node.searchWords)
		}
	}

	return result
}

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := new(Trie)
	trie.root = NewTrieNode()

	sort.Strings(products)

	for _, x := range products {
		trie.Insert(x)
	}

	return trie.Search(searchWord)
}

// Driver code
func main() {
	products := []string{"bat", "bag", "bassinet", "bread", "cable", "table", "tree", "tarp"}
	searchWordList := []string{"ba", "in", "ca", "t"}

	for i := range searchWordList {
		fmt.Printf("%d.\tProducts: ['%s']\n", i+1, strings.Join(products, "', '"))
		fmt.Printf("\tSearch keyword: \"%s\"\n", searchWordList[i])
		fmt.Printf("\tSuggested products: %s", doublyStringArrayToString(suggestedProducts(products, searchWordList[i])))
		fmt.Printf("\n%s\n", strings.Repeat("-", 93))
	}
}
