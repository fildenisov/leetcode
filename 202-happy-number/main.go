package main

import "fmt"

func main() {
	// fmt.Println(isHappy(3))
	fmt.Println(isHappy(19))
	// fmt.Println(isHappy(100))
}

func isHappy(n int) bool {
	if n < 0 {
		return false
	}

	res := -1
	for i := 0; i < 100; i++ {
		res = sqDigits(n)
		if res == 1 {
			return true
		}

		n = res
	}

	return false
}

func sqDigits(n int) int {
	res := 0
	tmp := 0
	sub := 0
	devider := 10

	fmt.Println("----------------")
	fmt.Printf("n = %v\n", n)

	for n > 0 {

		sub = n % devider
		if sub == 0 {
			devider *= 10
			continue
		}

		tmp = sub / (devider / 10)

		fmt.Println(sub)
		fmt.Println(tmp)

		if tmp == 0 {
			n -= tmp
			continue
		}

		res += tmp * tmp
		fmt.Printf("%v+", tmp*tmp)
		n -= sub
		devider *= 10
	}

	fmt.Printf("= %v\n", res)
	return res
}
