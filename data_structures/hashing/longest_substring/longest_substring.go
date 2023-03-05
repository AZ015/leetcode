package main

import "fmt"

func longestSubstring(str string) int {
	n, ans := len(str), 0
	m := make(map[int32]int)
	for i, j := 0, 0; j < n; j++ {
		ch := int32(str[j])
		if _, ok := m[ch]; ok {
			i = max(m[ch], i)
		}
		ans = max(ans, j-i+1)
		m[ch] = j + 1
	}

	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func main() {
	fmt.Println(longestSubstring("abcabcbb")) // 3
	fmt.Println(longestSubstring("bbbbb"))    // 1
	fmt.Println(longestSubstring("pwwkew"))   // 3
}
