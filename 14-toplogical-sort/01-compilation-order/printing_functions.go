package main

import "strconv"

// printGraph() is a helper function that prints the graph
func PrintGraph(graph map[rune][]rune) string {
	out := "["
	for index, value := range graph {
		if len(value) == 0 {
			out += "'" + string(index) + "'" + ":[], "
		} else {
			out += "'" + string(index) + "'" + ": ["
			for _, v := range value {
				out += "'" + string(v) + "'" + ", "
			}
			out = out[0 : len(out)-2]
			out += "], "
		}
	}
	out = out[0 : len(out)-2]
	out += "]"
	return out
}

// printInDegree() is a helper function that prints the inDegree map
func PrintInDegree(inDegree map[rune]int) string {
	out := "["
	for index, value := range inDegree {
		out += "'" + string(index) + "'" + ":" + strconv.Itoa(value) + ", "
	}
	out = out[0 : len(out)-2]
	out += "]"
	return out
}

/*
	runeArrayToString is used to convert an rune array to string.
	runeArrayToString is used as an helper function in a main function for
	printing purposes.
*/
func RuneArrayToString(char []rune) string {
	if len(char) == 0 {
		return "[]"
	}
	res := "["
	for _, num := range char {
		res += strconv.QuoteRune(num) + ", "
	}
	res += "]"
	return res[:len(res)-3] + "]"
}

/*
	runeDoubleArrayToString is used to convert a 2D rune array to string.
	runeDoubleArrayToString is used as an helper function in a main function for
	printing purposes.
*/
func RuneDoubleArrayToString(charB [][]rune) string {
	out := "["
	for _, char := range charB {
		out += "["
		for j, num := range char {
			if j < len(char)-1 {
				out += strconv.QuoteRune(num) + ", "
			} else {
				out += strconv.QuoteRune(num) + "], "
			}
		}
	}
	out2 := out[0 : len(out)-2]
	out2 += "]"
	return out2
}