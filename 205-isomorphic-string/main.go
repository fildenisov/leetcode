package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("add", "egg"))
}

func isIsomorphic(s string, t string) bool {
	dict := make(map[uint8]uint8)
	for i := 0; i < len(s); i ++ {
		dict[s[i]] = t[i]
	}


	for i := 0; i < len(s); i ++ {
		if dict[s[i]] != t[i] {
			return false
		}
	}

	return true
}
