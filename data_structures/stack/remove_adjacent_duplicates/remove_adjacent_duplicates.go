package main

import (
	"fmt"
	"strings"
)

func RemoveAdjacentDuplicates(str string) string {
	stack := make([]string, 0)

	for _, ch := range str {
		if len(stack) != 0 && stack[len(stack)-1] == string(ch) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(ch))
		}
	}

	return strings.Join(stack, "")
}

func main() {
	fmt.Println(RemoveAdjacentDuplicates("abbaca"))  // ca
	fmt.Println(RemoveAdjacentDuplicates("abac"))    // abac
	fmt.Println(RemoveAdjacentDuplicates("abcdefg")) // abcdefg
}
