package main

import (
	"fmt"
)

func main() {
	sb1 := []byte("hello")
	sb2 := []byte("buy")
	reverseString(sb1)
	fmt.Println(string(sb1))
	reverseString(sb2)
	fmt.Println(string(sb2))
}

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
