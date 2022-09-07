package main

import (
	"fmt"
	"sort"
)

/*

[ ][ ][ ]
[ ][ ][ ]
[ ][ ][ ]

0,0 + 1,1 + 2,2       3 & 3
0,2 + 1,1 + 2,0       3 & 3

0,0 + 0,1 + 0,2       0 & 3
1,0 + 1,1 + 1,2       3 & 3
2,0 + 2,1 + 2,2       6 & 3

0,0 + 1,0 + 2,0.      3 & 0
0,1 + 1,1 + 2,1       3 & 3
0,2 + 1,2 + 2,2       3 & 6


*/
var (
	winField = map[[3][2]int]struct{}{
		{{0,0}, {1,1}, {2,2}}: {},
		{{0,2}, {1,1}, {2,0}}: {},
		{{0,0}, {0,1}, {0,2}}: {},
		{{1,0}, {1,1}, {1,2}}: {},
		{{2,0}, {2,1}, {2,2}}: {},
		{{0,0}, {1,0}, {2,0}}: {},
		{{0,1}, {1,1}, {2,1}}: {},
		{{0,2}, {1,2}, {2,2}}: {},
	}
)

type sortable [][]int
func (s sortable) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		return s[i][1] < s[j][1]
	}

	return s[i][0] < s[j][0]
}
func (s sortable) Swap(i,j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortable) Len() int {
	return len(s)
}

func main() {
	// fmt.Println(tictactoe([][]int{{1,2},{1,0},{0,0},{0,1},{2,1}})) // pending
	fmt.Println(tictactoe([][]int{{0,0},{2,0},{1,1},{2,1},{2,2}})) // A

}

func tictactoe(moves [][]int) string {
	winner := "A"
	first := 0
	if len(moves) % 2 == 0 {
		winner = "B"
		first = 1
	}

	arr := sortable{{0,0},{0,0},{0,0}}
	checker := [3][2]int{}

	for x := first; x < len(moves); x += 2 {
		for y := first; y < len(moves); y += 2 {
			if x == y {
				continue
			}

			arr[0][0], arr[0][1] = moves[len(moves)-1][0], moves[len(moves)-1][1]
			arr[1][0], arr[1][1] = moves[x][0], moves[x][1]
			arr[2][0], arr[2][1] = moves[y][0], moves[y][1]

			sort.Sort(arr)
			checker = [3][2]int{{arr[0][0], arr[0][1]}, {arr[1][0], arr[1][1]}, {arr[2][0], arr[2][1]}}
			if _, ok := winField[checker]; ok {
				return winner
			}
 		}
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}

