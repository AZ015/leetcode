package main

import "fmt"

func occurrences(str string) bool {
	m := make(map[int32]int)

	for _, ch := range str {
		m[ch] += 1
	}

	chLen := m[int32(str[0])]

	for _, v := range m {
		if v != chLen {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(occurrences("abacbc"))
	fmt.Println(occurrences("aaabb"))
}
