package main

import "fmt"

func main() {
	//[1,2,3,0,0,0]
	//3
	//[2,5,6]
	//3
	//nums1 := []int{1,2,3,0,0,0}
	//m:= 3
	//nums2 := []int{2,5,6}
	//n := 3

	//nums1 := []int{0}
	//m := 0
	//nums2 := []int{1}
	//n := 1

	//[1,0]
	//1
	//[2]
	//1

	//nums1 := []int{1,0}
	//n := 1
	//nums2 := []int{2}
	//m := 1

	//[0,0,3,0,0,0,0,0,0]
	//3
	//[-1,1,1,1,2,3]
	//6

	nums1 := []int{0,0,3,0,0,0,0,0,0}
	m := 3
	nums2 := []int{-1,1,1,1,2,3}
	n := 6

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if n == 0 {
		return
	}

	i := 0
	cur := 0
	for j := 0; j < n; j ++ {
		for m != 0 && i < m+j && nums1[i] <= nums2[j] {
			i++
		}

		// shuffle forward
		cur = nums1[i]
		for x := i+1; x < m+n; x++ {
			nums1[x], cur = cur, nums1[x]
		}

		nums1[i] = nums2[j]
		i++
	}
}

//[1,2,3,0,0,0]    [1,2,2,3,0,0]

//[2,5,6]

