package main

import (
	"math"
	"strconv"
	"strings"
)

func alienOrder(words []string) string {
	adjList := make(map[rune][]rune)
	counts := make(map[rune]int)
	for _, word := range words {
		for _, c := range word {
			counts[c] = 0
			adjList[c] = make([]rune, 0)
		}
	}

    for i := 0; i < len(words) - 1; i++ {
        word1 := words[i]
        word2 := words[i + 1]

        if len(word1) > len(word2) && strings.Index(word1, word2) == 0 {
            return ""
        }

        m1 := []rune(word1)
        m2 := []rune(word2)
        length := int(math.Min(float64(len(m1)), float64(len(m2))))
        for j := 0; j < length; j++ {
            if m1[j] != m2[j] {
                present := false
                for _, value := range adjList[m1[j]] {
                    if value == m2[j] {
                        present = true
                        break
                    }
                }

                if !present {
                    adjList[m1[j]] = append(adjList[m1[j]], m2[j])
                    counts[m2[j]] = counts[m2[j]] + 1
                }
                break
            }
        }
    }

    var result string
    sourcesQueue := new(Queue)
    for char, itr := range counts {
        if itr == 0 {
            sourcesQueue.Enqueue(char)
        }
    }

    for !sourcesQueue.IsEmpty() {
        c := sourcesQueue.Dequeue().(rune)
        result += string(c)

        for _, next := range adjList[c] {
            counts[next] = counts[next] - 1
            if counts[next] == 0 {
                sourcesQueue.Enqueue(next)
            }
        }
    }

    if len(result) < len(counts) {
        return ""
    }

    return result
}

// printArrayWithMarkers() prints given array with markers on elements at given indices
func printArrayWithMarkers(words []string, indices []int) string {
	out := "["
	for index, word := range words {
		if index == indices[0] || index == indices[1] {
			out = out + "«" + word + "», "
		} else {
			out = out + word + ", "
		}
	}
	out = out[0 : len(out)-2]
	out += "]"
	return out
}

// printStringWithMarkers() prints the given string with markers on character at the given index
func printStringWithMarkers(word string, index int) string {
	return (word[0:index] + "«" + string(word[index:index+1]) + "»" + word[index+1:])
}

// printMap() is a helper function to print map as a string
func printMap(count map[rune]int) string {
	out := "["
	for i, v := range count {
		out = out + "'" + string(i) + "'" + ":" + strconv.Itoa(v) + ", "
	}
	out = out[0 : len(out)-2]
	out += "]"
	return out
}

// runeArrayToString() is helper function that converts rune array to a string for printing purposes
func runeArrayToString(runeArray []rune) string {
	if len(runeArray) == 0 {
		return "[]"
	}
	out := "["
	for _, v := range runeArray {
		out = out + "'" + string(v) + "', "
	}
	out = out[0 : len(out)-2]
	out += "]"
	return out
}
