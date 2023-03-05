package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[int32]int)

	for _, ch := range magazine {
		m[ch] += 1
	}

	for _, ch := range ransomNote {
		if _, ok := m[ch]; !ok {
			return false
		}

		m[ch] -= 1
		if m[ch] <= 0 {
			delete(m, ch)
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))    // false
	fmt.Println(canConstruct("aa", "ab"))  // false
	fmt.Println(canConstruct("aa", "aab")) // true
}
