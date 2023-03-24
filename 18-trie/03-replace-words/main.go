package main

import (
	"fmt"
	"strings"
)

func replaceWords(setence string, dictionary []string) string {
	trie := NewTrie()

	for _, prefix := range dictionary {
		trie.Insert(prefix)
	}

	newList := strings.Split(setence, " ")

	for i := range newList {
		newList[i] = trie.Replace(newList[i])
	}

	return strings.Join(newList, " ")
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

// Driver code
func main() {
	sentence := []string{"where there is a will there is a way",
		"the quick brown fox jumps over the lazy dog",
		"oops there is no matching word in this sentence",
		"i was born on twenty ninth february",
		"i dont know where you are but i will find you eventually"}
	dictionary := [][]string{{"wi", "wa", "w"},
		{"qui", "f", "la", "d"},
		{"oops", "there", "is", "no", "matching", "word", "in", "this", "sentence"},
		{"wa", "w", "a", "ty", "nint", "nin", "n", "feb", "februa", "f"},
		{"cool", "how", "sunday", "sun", "x"}}

	for i := range sentence {
		fmt.Printf("%d.\t Input sentence: \"%s\"\n", i+1, sentence[i])
		fmt.Printf("\t Dictionary words: %s\n", stringArrayToString(dictionary[i]))
		fmt.Printf("\t After replacing words: \"%s\"\n", replaceWords(sentence[i], dictionary[i]))
		fmt.Printf("%s\n", strings.Repeat("-", 95))
	}
}
