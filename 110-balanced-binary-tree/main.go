package main

import "fmt"

func main() {

	n15 := TreeNode{Val: 15}
	n7 := TreeNode{Val: 7}
	n20 := TreeNode{Val: 20, Left: &n15, Right: &n7}

	n9 := TreeNode{Val: 9}

	n3 := TreeNode{Val: 3, Left: &n9, Right: &n20}

	fmt.Println(isBalanced(&n3))

	n4_1 := TreeNode{Val: 4}
	n4_2 := TreeNode{Val: 4}

	n3_1 := TreeNode{Val: 3, Left: &n4_1, Right: &n4_2}
	n3_2 := TreeNode{Val: 3}

	n2_1 := TreeNode{Val: 2, Left: &n3_1, Right: &n3_2}
	n2_2 := TreeNode{Val: 2}

	n1 := TreeNode{Val: 1, Left: &n2_1, Right: &n2_2}

	fmt.Println(isBalanced(&n1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	b := 0

	if root.Left != nil {
		b++

		add := false
		addMore := false
		if root.Left.Left != nil {
			add = true

			if root.Left.Left.Left != nil || root.Left.Left.Right != nil {
				addMore = true
			}
		}

		if root.Left.Right != nil {
			add = true

			if root.Left.Right.Left != nil || root.Left.Right.Right != nil {
				addMore = true
			}
		}

		if add {
			b++
		}

		if addMore {
			b++
		}

	}

	if root.Right != nil {
		b--

		sub := false
		subMore := false
		if root.Right.Left != nil {
			sub = true

			if root.Right.Left.Left != nil || root.Right.Left.Right != nil {
				subMore = true
			}
		}

		if root.Right.Right != nil {
			sub = true

			if root.Right.Right.Left != nil || root.Right.Right.Right != nil {
				subMore = true
			}
		}

		if sub {
			b--
		}

		if subMore {
			b--
		}

	}

	if b < -1 || b > 1 {
		return false
	}

	if root.Left != nil && !isBalanced(root.Left) {
		return false
	}

	if root.Right != nil && !isBalanced(root.Right) {
		return false
	}

	return true
}
