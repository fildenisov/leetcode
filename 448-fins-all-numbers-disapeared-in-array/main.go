package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findDisappearedNumbers([]int{4,3,2,7,8,2,3,1}))
}

func findDisappearedNumbers(nums []int) []int {
	sort.Ints(nums)

	fmt.Println(nums)

	shift := 0
	l := 0
	for i := 0; i < len(nums); i ++ {
		if i > 0 && nums[i] == nums[i-1] {
			shift++
			continue
		}
		// TODO: trick here
		if i+1-shift != nums[i] {
			nums[l] = i+1-shift
			l++
		}
	}

	return nums[:l]
}
