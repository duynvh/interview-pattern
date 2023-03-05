package main

type Stack struct {
	store []rune
}

func (s *Stack) remove(ch rune) bool {
    if len(s.store) == 0 {
        return false
    }

    switch ch {
    case ')':
        if s.store[len(s.store) - 1] != '(' {
            return false
        }
    case ']':
        if s.store[len(s.store) - 1] != '[' {
            return false
        }
    case '}':
        if s.store[len(s.store) - 1] != '{' {
            return false
        }
    }

    s.store = s.store[:len(s.store) - 1]
    return true
}

func validParentheses(s string) bool {
	stack := Stack{
        store: []rune{},
    }

    for _, ch := range s {
        switch{
        case ch == '(' || ch == '[' || ch == '{':
            stack.store = append(stack.store, ch)
        default:
            if ok := stack.remove(ch); !ok {
                return false
            }
        }
    }
	return len(stack.store) == 0
}