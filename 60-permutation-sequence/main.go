package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("RESULT = " + getPermutation(3, 3)) // 213
	fmt.Println("RESULT = " + getPermutation(4, 9)) // 2314
	fmt.Println("RESULT = " + getPermutation(3, 1)) // 123
	fmt.Println("RESULT = " + getPermutation(3, 2)) // 132
	fmt.Println("RESULT = " + getPermutation(3, 5)) // 312
	fmt.Println("RESULT = " + getPermutation(3, 6)) // 321
}

func dbg(s string) {
	// fmt.Print(s)
}

var (
	factorials = [10]int{0, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
)

func getPermutation(n int, k int) string {
	sb := strings.Builder{}
	sb.Grow(n)

	opts := make([]uint8, n)
	max := uint8(n)
	for i := uint8(0); i < max; i++ {
		opts[i] = 49 + i
	}

	ss, i := 0, 0
	for n != 1 {
		dbg(fmt.Sprintf("opts = %v\n", opts))
		dbg(fmt.Sprintf("n = %v, k = %v\n", n, k))
		ss = factorials[n] / n
		dbg(fmt.Sprintf("ss = %v\n", ss))
		i = k / ss
		if k%ss > 0 {
			i++
		}

		dbg(fmt.Sprintf("i = %v\n", i))
		sb.WriteByte(opts[i-1])
		k = k - (i-1)*ss
		opts = append(opts[:i-1], opts[i:]...)
		n--
		dbg(fmt.Sprintf("res = %v\n", sb.String()))
		dbg(fmt.Sprintln("------------"))
	}
	sb.WriteByte(opts[0])
	return sb.String()
}
