package simplify_path

import (
	"fmt"
	. "stack"
	"strings"
)

func simplifyPath(path string) string {
	stack := New()

	for _, ch := range strings.Split(path, "/") {
		if string(ch) == ".." {
			stack.Pop()
		} else if string(ch) == "." || string(ch) == "" {
			continue
		} else {
			stack.Push(ch)
		}
	}

	return fmt.Sprintf("/%v", stack.ToString("/"))
}

func main() {
	fmt.Println(simplifyPath("/home/"))      // /home
	fmt.Println(simplifyPath("/../"))        // /
	fmt.Println(simplifyPath("/home//foo/")) // /home/foo
}
