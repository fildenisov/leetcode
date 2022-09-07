package main

import "fmt"

func main() {
	fmt.Println(addDigits(546123))
}

func addDigits(num int) int {
	var dev, numOrig, toAdd int
	for num > 9 {
		fmt.Printf("NUM = %v\n", num)
		dev = 10
		numOrig = num
		num = 0
		for numOrig > 0 {
			if dev < 0 {
				panic("oveflow")
			}

			toAdd = numOrig % dev
			numOrig -= toAdd
			toAdd /= dev / 10
			fmt.Printf("toAdd = %v\n", toAdd)
			num += toAdd
			// fmt.Printf("dev = %v, numOrig = %v, num = %v\n", dev, numOrig, num)
			dev *= 10
		}
	}
	return num
}

/*
   546123 = 5 + 4 + 6 + 1 + 2 + 3 = 21
    dev = 10
    546123 % 10 = 3
    orig = 546123 - 3 = 546120
    3 / 1 = 3

    dev = 100
    546120 % 100 = 20
    orig = 546120 - 20 = 546100
    20 / (100/10) = 2



    n % 100 = 20 / (dev/10) =

   21 = 2 + 1 = 3

*/
