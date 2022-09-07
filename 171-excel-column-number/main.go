package main

import (
	"fmt"
	"math"
)

const (
	shift = 64
	alphaSize float64 = 26
)

func main() {
	fmt.Println(titleToNumber("AB")) // 28
	fmt.Println(titleToNumber("A")) // 1
	fmt.Println(titleToNumber("Z")) // 26
	fmt.Println(titleToNumber("ZY")) // 701
}

func titleToNumber(columnTitle string) int {
	b := []byte(columnTitle)

	res := 0
	med := 0
	var power float64
	for i := len(b)-1; i >= 0; i -- {
		med = int(b[i]) - shift
		if power != 0 {
			med *= int(math.Pow(alphaSize, power))
		}
		res += med
		power++
	}
	return res
}
