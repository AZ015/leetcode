package graphs

import "fmt"

type Graph struct {
	neighbors map[int][]int
}

func NewDirectlyGraph(edges [][]int) *Graph {
	g := make(map[int][]int)

	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
	}

	return &Graph{
		neighbors: g,
	}
}

func NewUnDirectlyGraph(edges [][]int) *Graph {
	g := make(map[int][]int)

	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	return &Graph{
		neighbors: g,
	}
}

func main() {
	directly := NewDirectlyGraph([][]int{{0, 1}, {1, 2}, {2, 0}, {2, 3}})
	fmt.Println(directly)
	unDirectly := NewUnDirectlyGraph([][]int{{0, 1}, {1, 2}, {2, 0}, {2, 3}})
	fmt.Println(unDirectly)
}
