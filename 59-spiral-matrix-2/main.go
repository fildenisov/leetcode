package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(4))
}

var (
	dirs = [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
)

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	x, y := 0, 0
	curDir := 0

	// dir = 0, +1
	// x, y = 0, 0
	// x, y = 0, 1
	// x, y = 0, 2

	//dir = +1, 0
	// x, y = 1, 2
	// x, y = 2, 2

	//dir = 0, -1
	// x, y = 2, 1
	// x, y = 2, 0

	//dir = -1, 0
	// x, y = 1, 0

	//dir = 0, +1
	// x, y = 1, 1

	// if next where you r going is not zero - switch direction
	nX, nY := 0, 0
	switchDir := func() {
		if curDir == 3 {
			curDir = 0
		} else {
			curDir++
		}
		nX, nY = x+dirs[curDir][0], y+dirs[curDir][1]
	}

	total := n * n

	for i := 1; i <= total; i++ {
		res[x][y] = i
		if i == total {
			break
		}
		// fmt.Println(res[x][y])
		nX, nY = x+dirs[curDir][0], y+dirs[curDir][1]
		if nX < 0 || nY < 0 || nX == n || nY == n {
			switchDir()
		}
		if res[nX][nY] != 0 {
			switchDir()
		}

		x, y = nX, nY
	}

	return res
}
