package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(summaryRanges([]int{0,1,2,4,5,7})) // ["0->2","4->5","7"]
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	res := make([]string, 0, 1)

	start := nums[0]
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(nums[0]))

	for i := 1; i < len(nums); i ++ {
		if nums[i] > nums[i-1] + 1 {
			if nums[i-1] != start {
				sb.WriteString("->")
				sb.WriteString(strconv.Itoa(nums[i-1]))
			}

			res = append(res, sb.String())
			sb.Reset()
			start = nums[i]
			sb.WriteString(strconv.Itoa(nums[i]))
		}
	}

	if start != nums[len(nums)-1] {
		sb.WriteString("->")
		sb.WriteString(strconv.Itoa(nums[len(nums)-1]))
	}

	res = append(res, sb.String())
	return res
}
