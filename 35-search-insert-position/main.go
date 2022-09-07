package main

import "fmt"

func main() {
	// fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0)) // 0
	fmt.Println(searchInsert([]int{1, 3, 5}, 4)) // 2
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		}
		return 1
	}

	i := len(nums) / 2 // 0
	left := 0
	right := len(nums) // 1

	stopper := 0

	for {
		fmt.Printf("i = %v | left = %v | right = %v\n", i, left, right)
		if left == right || left+1 == right || right+1 == left {
			if nums[left] >= target {
				return left
			}
			return left + 1
		}
		switch {
		case nums[i] > target:
			right = i
			i = (i - left - 1) / 2
		case nums[i] < target:
			left = i
			i = left + (right-i)/2
		default:
			return i
		}
		stopper++
		if stopper == 5 {
			return -666
		}
	}
}
