package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

type TimeStamp struct {
	valuesMap    map[string][]string
	timestampMap map[string][]int
}

func Constructor() TimeStamp {
	return TimeStamp{
		valuesMap:    make(map[string][]string),
		timestampMap: make(map[string][]int),
	}
}

func (this *TimeStamp) SetValue(key string, value string, timeStamp int) bool {
	saved := false
	if _, ok := this.valuesMap[key]; ok {
		if value != this.valuesMap[key][len(this.valuesMap[key])-1] {
			this.valuesMap[key] = append(this.valuesMap[key], value)
			this.timestampMap[key] = append(this.timestampMap[key], timeStamp)
			saved = true
		}
	} else {
		this.valuesMap[key] = append(this.valuesMap[key], value)
		this.timestampMap[key] = append(this.timestampMap[key], timeStamp)
		saved = true
	}

	return saved
}

func (this *TimeStamp) GetValue(key string, timestamp int) string {
	if _, ok := this.valuesMap[key]; !ok {
		return ""
	} else {
		index := sort.Search(len(this.timestampMap[key]), func(i int) bool {
			return this.timestampMap[key][i] > timestamp
		}) - 1

		if index > -1 {
			return this.valuesMap[key][index]
		}

		return ""
	}
}

// main is used for driver code
func main() {
	ts := Constructor()
	keyInputList := []string{"course", "course", "course", "course", "course"}
	valueInputList := []string{"OOP", "PF", "OS", "ALGO", "DB"}
	timeStampInputList := []int{3, 4, 5, 7, 8}

	for i := range keyInputList {
		fmt.Printf("%d.\tAdd value: (\"%s\", \"%s\", %d)\n", i+1, keyInputList[i], valueInputList[i], timeStampInputList[i])
		ts.SetValue(keyInputList[i], valueInputList[i], timeStampInputList[i])
		randomValue := rand.Intn(10)
		fmt.Print("\n\tGet value for:\n")
		fmt.Printf("\t\tKey = \"course\"  \n\t\tTimestamp = %d\n", randomValue)
		fmt.Printf("\n\tReturned value = \"%s\"\n", ts.GetValue("course", randomValue))
		fmt.Printf("%s\n", strings.Repeat("-", 95))
	}
}
