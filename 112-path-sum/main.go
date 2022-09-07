package main

import "fmt"

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

  func main() {
  	l := TreeNode{
  		Val: 2,
	}

	r := TreeNode{
		Val:   1,
		Left:  &l,
	}

	fmt.Println(hasPathSum(&r, 2))
  }

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil ||
		(root.Val == targetSum && (root.Right != nil || root.Left != nil)) {
		return false
	}

	return findPathSum(root, targetSum, 0)
}

func findPathSum(node *TreeNode, target, current int) bool {
	if node != nil {
		current+= node.Val
	}

	if current == target {
		return true
	}

	if node.Left != nil && findPathSum(node.Left, target, current) {
		return true
	}

	if node.Right != nil && findPathSum(node.Right, target, current) {
		return true
	}

	return false
}
