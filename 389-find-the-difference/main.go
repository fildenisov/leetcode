package main

import "fmt"

func main() {
	fmt.Println(string(findTheDifference("abcd", "abcde")))
	fmt.Println(findTheDifference("abcd", "abcde"))
}

func findTheDifference(s string, t string) byte {
	var res byte
	for i := 0; i < len(t); i ++ {
		res += t[i]
		if i != len(s) {
			res -= s[i]
		}
	}
	return res
}
