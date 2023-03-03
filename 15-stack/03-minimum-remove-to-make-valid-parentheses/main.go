package main

import (
	"fmt"
	"strings"
)

func minRemoveParentheses(s string) string {
	delimiters := Stack{}
	builder := s

	for i, val := range s {
		if len(delimiters) > 0 && delimiters.Top().value == '(' && val == ')' {
			delimiters, _ = delimiters.Pop()
		} else if val == '(' || val == ')' {
			delimiters = delimiters.Push(Set{val, i})
		}
	}

	poppedVal := Set{}
	for !delimiters.Empty() {
		delimiters, poppedVal = delimiters.Pop()
		index := poppedVal.index
		builder = builder[0:index] + builder[index+1:]
	}

	return builder
}

// Driver code
func main() {
	inputs := []string{
        "a)di(o)qw(",
        // "ar)ab(abc)abd(", 
        // "a)rt)lm(ikgh)", 
        // "aq)xy())qf(a(ba)q)", 
        // "(aw))kk())(w(aa)(bv(wt)r)", 
        // "(qi)(kl)((y(yt))(r(q(g)s)",
    }
	for i, input := range inputs {
		fmt.Printf("%d.\tInput: \"%s\"\n", i+1, input)
		fmt.Printf("\tValid parentheses, after minimum removal: \"%s\"\n", minRemoveParentheses(input))
		fmt.Printf("%s\n", strings.Repeat("-", 93))
	}
}
