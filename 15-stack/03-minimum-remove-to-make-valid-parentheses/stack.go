package main

type Set struct {
    value rune
    index int
}

type Stack []Set

func (s Stack) Push(v Set) Stack {
    return append(s, v)
}

func (s Stack) Pop() (Stack, Set) {
    l := len(s)
    if l == 0 {
        return s, Set{0, 0}
    }

    return s[:l-1], s[l-1]
}

func (s Stack) Top() Set {
    l := len(s)
    return s[l-1]
}

func (s Stack) Empty() bool {
    return len(s) == 0
}