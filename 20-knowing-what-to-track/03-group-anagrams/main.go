package main

type CharCount [26]int

func groupAnagrams(strs []string) [][]string {
    var result [][]string
    hash := make(map[CharCount][]string)

    for _, word := range strs {
        charCount := CharCount{}
        for i := 0; i < len(word); i++ {
            charCount[word[i] - 'a']++
        }

        hash[charCount] = append(hash[charCount], word)
    }

    for _, value := range hash {
        result = append(result, value)
    }

    return result
}