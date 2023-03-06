package main

import "fmt"

func isValid(str string) bool {
	stack := make([]string, 0)
	matching := map[string]string{"(": ")", "{": "}", "[": "]"}

	for _, par := range str {
		if _, ok := matching[string(par)]; ok {
			stack = append(stack, string(par))
		} else {
			if len(stack) == 0 {
				return false
			}

			previous := stack[len(stack)-1]
			if matching[previous] != string(par) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("({})"))
	fmt.Println(isValid("(){}[]"))
	fmt.Println(isValid("(]"))
}
