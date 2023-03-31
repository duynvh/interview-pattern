package main

func longestPalindrome(s string) int {
    if len(s) < 2 {
        return len(s)
    }

    m := make(map[byte]int)
    oddCount := 0

    for i := 0; i < len(s); i++ {
        m[s[i]]++
    }

    for _, val := range m {
        if val % 2 != 0 {
            oddCount++
        }
    }

    if oddCount == 0 {
        return len(s)
    }

    return len(s) + 1 - oddCount
}