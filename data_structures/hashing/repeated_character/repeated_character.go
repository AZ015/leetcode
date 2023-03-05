package main

import (
	"fmt"
)

func repeatedCharacter(str string) string {
	m := make(map[int32]struct{})

	for _, char := range str {
		if _, ok := m[char]; ok {
			return string(char)
		}
		m[char] = struct{}{}
	}

	return ""
}

func main() {
	fmt.Println(repeatedCharacter("abcdeda"))
}
