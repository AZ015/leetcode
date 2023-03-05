package main

import (
	"fmt"
	"strings"
)

func checkIsPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	m := make(map[string]struct{})

	for _, ch := range sentence {
		chToLower := strings.ToLower(string(ch))
		m[chToLower] = struct{}{}
	}

	if len(m) < 26 {
		return false
	}

	return true
}

func main() {
	fmt.Println(checkIsPangram("leetcode"))                            // false
	fmt.Println(checkIsPangram("thequickbrownfoxjumpsoverthelazydog")) // true
}
