package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("11", "1"))      // 100
	fmt.Println(addBinary("0", "0"))       // 0
	fmt.Println(addBinary("1", "1"))       // 10
	fmt.Println(addBinary("1111", "1111")) // 10000
	fmt.Println(addBinary("1010", "1011")) // 10101
}

const (
	zero = uint8(48)
	one  = uint8(49)
)

func addBinary(a string, b string) string {
	add := false
	maxLen := len(a)
	if len(b) > len(a) {
		maxLen = len(b)
	}

	res := make([]byte, maxLen+1)
	i := 0
	var cur uint8
	sum := 0
	for {
		i++
		if i > len(a) && i > len(b) {
			break
		}

		sum = 0
		if i <= len(a) && a[len(a)-i] == one {
			sum++
		}

		if i <= len(b) && b[len(b)-i] == one {
			sum++
		}

		if add {
			sum++
		}

		switch sum {
		case 3:
			add = true
			cur = one
		case 2:
			add = true
			cur = zero
		case 1:
			add = false
			cur = one
		default:
			add = false
			cur = zero
		}

		char := "0"
		if cur == one {
			char = "1"
		}
		fmt.Printf("res[%v] = %v | add = %v\n", len(res)-i, char, add)
		res[len(res)-i] = cur
	}

	if add {
		res[0] = one
	} else {
		res = res[1:]
	}

	return string(res)
}
