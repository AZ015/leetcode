package main

import (
	"fmt"
)

func build(str string) string {
	res := make([]int32, 0)
	for _, ch := range str {
		if string(ch) != "#" {
			res = append(res, ch)
		} else {
			res = res[:len(res)-1]
		}
	}

	return string(res)
}

func backspaceCompare(str string, t string) bool {
	return build(str) == build(t)
}

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c")) // true
	fmt.Println(backspaceCompare("ab#d", "ad#c")) // false
}
