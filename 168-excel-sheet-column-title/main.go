package main

import (
	"fmt"
	"strings"
)

//const (
//	shift = 64
//	alphaSize float64 = 26
//)

func main() {
	fmt.Println(convertToTitle(26)) // Z
	// fmt.Println(convertToTitle(27)) // AA
	// fmt.Println(convertToTitle(28)) // AB
	// fmt.Println(convertToTitle(701)) // ZY
	// fmt.Println(convertToTitle(54)) // BB
}

/*
// 54
// 54 % 26 = 2
// res = B
// 54 - 2 = 52
// 52 % 676 = 52

26
26%26 = 0
26 - 26//26 = 25
Z

al1 = 10
al2 = 26

1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30
A B C D E F G H I J  K  L  M  N  O  P  Q  R  S  T  U  V  W  X  Y  Z  AA AB AC AD AE

cap = 26

Z = 26 % cap+1 = 26

AA = 27 % cap+1 = 0 => C27%cap = 1, x -= 1



BB = 54 % CAP+1 = 0 =>

 */

const (
	shift int32 = 64
	alphaSize = 26
)

func convertToTitle(columnNumber int) string {
	sb := strings.Builder{}

	var cur int32
	for {
		if columnNumber <= 0 {
			break
		}


	}

	s := []byte(sb.String())
	for i := 0; i < len(s) / 2; i ++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}

	return string(s)
}

//
//func convertToTitle(columnNumber int) string {
//	s := strconv.FormatInt(int64(columnNumber), 26)
//	fmt.Println(s)
//	sb := strings.Builder{}
//	sb.Grow(len(s))
//
//	var res rune
//	for _, c := range s {
//		switch c {
//		case 49: res = 'A'
//		case 50: res = 'B'
//		case 51: res = 'C'
//		case 52: res = 'D'
//		case 53: res = 'E'
//		case 54: res = 'F'
//		case 55: res = 'G'
//		case 56: res = 'H'
//		case 57: res = 'I'
//		case 58: res = 'J'
//		case 59: res = 'K'
//		case 60: res = 'L'
//		case 61: res = 'M'
//		case 62: res = 'N'
//		case 63: res = 'O'
//		case 64: res = 'P'
//		case 65: res = 'Q'
//		case 66: res = 'R'
//		case 67: res = 'S'
//		case 68: res = 'T'
//		case 69: res = 'U'
//		case 70: res = 'V'
//		case 71: res = 'W'
//		case 72: res = 'X'
//		case 73: res = 'Y'
//		case 74: res = 'Z'
//		}
//		sb.WriteRune(res)
//	}
//	return sb.String()
//}