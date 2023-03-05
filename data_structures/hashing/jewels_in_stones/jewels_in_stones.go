package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	m := make(map[int32]int)

	for _, jewel := range jewels {
		m[jewel] += 1
	}

	numJewels := 0

	for _, stone := range stones {
		if _, ok := m[stone]; ok {
			numJewels += 1
		}
	}

	return numJewels
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb")) // 3
	fmt.Println(numJewelsInStones("z", "ZZ"))       // 0
}
