package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}

func firstUniqChar(s string) int {
	broken := false
	for i := 0; i < len(s); i++ {
		broken = false
		for j := 0; j < len(s); j ++ {
			if i == j {
				continue
			}
			if s[i] == s[j] {
				broken = true
				break
			}
		}
		if !broken {
			return i
		}
	}

	return -1
}
