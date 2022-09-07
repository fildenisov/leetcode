package main

import "fmt"

func main() {
	getRow(5)
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	res := make([]int, rowIndex+1)
	res[0] = 1
	res[1] = 1

	if rowIndex == 1 {
		return  res
	}

	val := 1

	for ri := 2; ri <= rowIndex; ri++{
		func(){
			for i := 1; i < ri; i ++ {
				val = res[i-1] + res[i]

				defer func(i, val int) {
					res[i] = val
				}(i, val)
			}
			res[ri] = 1
		}()
		fmt.Println(res)
	}

	return res
}

/*
[1 0 0 0 0]
[1 1 0 0 0]
...
[1 2 1 0 0]
[1 3 3 1 0]
[1 4 6 4 1]
 */










func getRowOld(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i ++ {
		res[i] = 1
	}

	if rowIndex <= 1 {
		return res
	}

	val := 1
	mid := (rowIndex + 1)/2
	fmt.Printf("mid=%v\n", mid)
	for x := 1; x < rowIndex; x++{
		func(){
			for i := 1; i < rowIndex; i ++ {
				val = res[i-1] + res[i+1]

				defer func(i, val int) {
					res[i] = val
				}(i, val)
			}
		}()
		fmt.Println(res)
	}

	/*
	[1 2 2 2 1]
	[1 3 4 3 1]
	[1 5 6 5 1]
	*/

	// [1 1 1 1 1]
	// [1 2 2 2 1]
	// [1 3 4 3 1]
	// [1 4 6 4 1]

	return res
}
