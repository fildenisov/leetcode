package main

import "sort"

const (
	x = 'X'
	o = 'O'
)

var (
	winField = map[[3][2]int]struct{}{
		{{0, 0}, {1, 1}, {2, 2}}: {},
		{{0, 2}, {1, 1}, {2, 0}}: {},
		{{0, 0}, {0, 1}, {0, 2}}: {},
		{{1, 0}, {1, 1}, {1, 2}}: {},
		{{2, 0}, {2, 1}, {2, 2}}: {},
		{{0, 0}, {1, 0}, {2, 0}}: {},
		{{0, 1}, {1, 1}, {2, 1}}: {},
		{{0, 2}, {1, 2}, {2, 2}}: {},
	}
)

type sortable [][]int

func (s sortable) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		return s[i][1] < s[j][1]
	}

	return s[i][0] < s[j][0]
}
func (s sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortable) Len() int {
	return len(s)
}

func validTicTacToe(board []string) bool {
	xMoves, oMoves := [][2]int{}, [][2]int{}
	for i := 0; i < len(board); i++ {
		for j, r := range board[i] {
			switch r {
			case x:
				xMoves = append(xMoves, [2]int{i, j})
			case o:
				oMoves = append(oMoves, [2]int{i, j})
			}
		}
	}

	if len(oMoves) > len(xMoves) || len(xMoves) > len(oMoves)+1 {
		return false
	}

	//check if only 1 win combo is on the board
	arr := sortable{{0, 0}, {0, 0}, {0, 0}}
	checker := [3][2]int{}
	gotWinnerX := false
	for a := 0; a < len(xMoves); a++ {
		for b := 0; b < len(xMoves); b++ {
			for c := 0; c < len(xMoves); c++ {
				if a == b || b == c || a == c {
					continue
				}

				arr[0][0], arr[0][1] = xMoves[a][0], xMoves[a][1]
				arr[1][0], arr[1][1] = xMoves[b][0], xMoves[b][1]
				arr[2][0], arr[2][1] = xMoves[c][0], xMoves[c][1]

				sort.Sort(arr)
				checker = [3][2]int{{arr[0][0], arr[0][1]}, {arr[1][0], arr[1][1]}, {arr[2][0], arr[2][1]}}
				if _, ok := winField[checker]; ok {
					gotWinnerX = true
				}
			}
		}
	}

	gotWinnerO := false
	for a := 0; a < len(oMoves); a++ {
		for b := 0; b < len(oMoves); b++ {
			for c := 0; c < len(oMoves); c++ {
				if a == b || b == c || a == c {
					continue
				}

				arr[0][0], arr[0][1] = oMoves[a][0], oMoves[a][1]
				arr[1][0], arr[1][1] = oMoves[b][0], oMoves[b][1]
				arr[2][0], arr[2][1] = oMoves[c][0], oMoves[c][1]

				sort.Sort(arr)
				checker = [3][2]int{{arr[0][0], arr[0][1]}, {arr[1][0], arr[1][1]}, {arr[2][0], arr[2][1]}}
				if _, ok := winField[checker]; ok {
					if gotWinnerX {
						return false
					}
					gotWinnerO = true
				}
			}
		}
	}

	if gotWinnerX && len(oMoves) == len(xMoves) {
		return false
	}

	if gotWinnerO && len(oMoves) != len(xMoves) {
		return false
	}

	return true
}

/*

["XXX","OOX","OOX"]

[X][X][X]
[O][O][X]
[O][O][X]

["XXX","XOO","OO "]

[X][X][X]
[O][O][X]
[O][O][ ]

*/
