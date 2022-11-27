package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayPairSum([]int{4, 2, 3, 1}))
	fmt.Println(arrayPairSum([]int{6, 2, 6, 5, 1, 2}))
	fmt.Println(arrayPairSum([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(arrayPairSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

func arrayPairSum(nums []int) int {
	var res int
	sort.Ints(nums)
	for i := len(nums) - 2; i >= 0; i -= 2 {
		res += nums[i]
	}

	return res
}
