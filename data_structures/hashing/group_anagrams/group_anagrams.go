package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(str []string) [][]string {
	groups := make(map[string][]string)

	for _, word := range str {
		key := SortString(word)
		groups[key] = append(groups[key], word)
	}

	res := make([][]string, 0)

	for _, words := range groups {
		res = append(res, words)
	}

	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
