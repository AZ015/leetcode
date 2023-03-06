package main

import (
	"fmt"
	s "stack"
	"unicode"
)

func makeGood(str string) string {
	stack := s.New()

	for _, ch := range str {
		if unicode.IsLower(ch) {
			stack.Push(string(ch))

		} else if unicode.IsUpper(ch) && string(unicode.ToLower(ch)) == stack.Peek() {
			stack.Pop()
		}
	}

	return stack.ToString("")
}

func main() {
	fmt.Println(makeGood("abBAcC"))     // ""
	fmt.Println(makeGood("leEeetcode")) // "leetcode"
	fmt.Println(makeGood("s"))          // "s"
}
