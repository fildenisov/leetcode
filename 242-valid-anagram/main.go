package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {return false}
	m := [26]uint16{}
	var i int
	for ; i < len(s); i ++ {
		m[s[i]-97]++
		m[t[i]-97]--
	}

	i = 0
	for ; i < len(m); i++ {
		if m[i] != 0 {return false}
	}
	return true
}
