package main

import (
	"fmt"
	"sort"
)

func findWinners(matches [][]int) [][]int {
	wm := make(map[int]int)
	lm := make(map[int]int)

	for _, match := range matches {
		winner, lost := match[0], match[1]
		wm[winner] += 1
		lm[lost] += 1
	}

	res := make([][]int, 2)

	for player := range wm {
		if _, ok := lm[player]; !ok {
			res[0] = append(res[0], player)
		}
	}

	for player, lostCount := range lm {
		if lostCount == 1 {
			res[1] = append(res[1], player)
		}
	}

	sort.Ints(res[0])
	sort.Ints(res[1])

	return res
}

func findWinnersSecondSolution(matches [][]int) [][]int {
	lossesCount := make(map[int]int)

	for _, match := range matches {
		winner, loser := match[0], match[1]

		if _, ok := lossesCount[winner]; !ok {
			lossesCount[winner] = 0
		}
		lossesCount[loser] += 1
	}

	zeroLost := make([]int, 0)
	oneLost := make([]int, 0)

	for player, count := range lossesCount {
		if count == 0 {
			zeroLost = append(zeroLost, player)
		}

		if count == 1 {
			oneLost = append(oneLost, player)
		}
	}

	sort.Ints(zeroLost)
	sort.Ints(oneLost)

	return [][]int{zeroLost, oneLost}
}

func main() {
	fmt.Println(findWinnersSecondSolution([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}))
	fmt.Println(findWinnersSecondSolution([][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}))
}

// not lost any match
// lost exactly one match
