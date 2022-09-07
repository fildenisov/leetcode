package main

import "fmt"

func main() {
	// [1, 4, 3, 2]
	// expected [1,4,2,3]

	n2 := &TreeNode{
		Val:   2,
	}

	n3 := &TreeNode{
		Val:   3,
	}

	n4 :=&TreeNode{
		Val:   4,
		Left:  n2,
	}

	n1 := &TreeNode{
		Val:   1,
		Left:  n4,
		Right: n3,
	}

	fmt.Println(preorderTraversal(n1))
}

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

type Task struct {
	Val *TreeNode
	Next *Task
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0, 1)

	task := &Task{
		Val: root,
	}

	for task != nil {
		res = append(res, task.Val.Val)

		if task.Val.Right != nil {
			task.Next = &Task{
				Val: task.Val.Right,
				Next: task.Next,
			}
		}

		if task.Val.Left != nil {
			task.Next = &Task{
				Val: task.Val.Left,
				Next: task.Next,
			}
		}

		task = task.Next
	}

	return res
}
