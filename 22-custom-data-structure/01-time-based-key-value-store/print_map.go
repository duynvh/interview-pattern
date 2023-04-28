package main

import (
	"fmt"
	"strings"
)

// stringStringMapToString is used to conver map[string][]string type
// map to string
func stringStringMapToString(hashMap map[string][]string) string {
	res := "{"
	for key, data := range hashMap {
		res += "\"" + key + "\": "
		out := strings.Join(data, "\", \"")
		out = "[\"" + out[0:1] + out[1:len(out)] + "\"]"
		res += out
	}
	res += "}"
	return res
}

// stringStringMapToString is used to conver map[string][]int type
// map to string
func stringIntMapToString(hashMap map[string][]int) string {
	res := "{"
	for key, data := range hashMap {
		res += "\"" + key + "\": "
		out := strings.Replace(fmt.Sprint(data), " ", ", ", -1)
		res += out
	}
	res += "}"
	return res
}
