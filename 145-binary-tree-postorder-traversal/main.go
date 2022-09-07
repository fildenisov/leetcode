package main

import "fmt"

func main() {
	// [3,1,2]

	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{
		Val:   3,
		Left:  n1,
		Right: n2,
	}

	fmt.Println(postorderTraversal(n3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Task struct {
	Val  *TreeNode
	Next *Task
	Prev *Task
}

func postorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return []int{}
	}

	res = make([]int, 0, 1)

	task := &Task{
		Val: root,
	}

	for task != nil {
		defer func(val int){
			res = append(res, val)
		}(task.Val.Val)

		if task.Val.Left != nil {
			task.Next = &Task{
				Val: task.Val.Left,
				Next: task.Next,
			}
		}

		if task.Val.Right != nil {
			task.Next = &Task{
				Val: task.Val.Right,
				Next: task.Next,
			}
		}

		task = task.Next
	}

	return res
}

//func postorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return []int{}
//	}
//
//	size := 0
//	task := &Task{
//		Val: root,
//	}
//	var tail *Task
//
//	for task != nil {
//		size++
//
//		if task.Val.Right != nil {
//			task.Next = &Task{
//				Val:  task.Val.Right,
//				Next: task.Next,
//				Prev: task,
//			}
//		}
//
//		if task.Val.Left != nil {
//			nt := &Task{
//				Val:  task.Val.Left,
//				Next: task.Next,
//				Prev: task,
//			}
//			if task.Next != nil {
//				nt.Next = task.Next.Next
//				nt.Prev = task.Next
//			}
//			task.Next = nt
//		}
//
//		tail = task
//		for tail.Next != nil {
//			tail = tail.Next
//		}
//
//		task = task.Next
//	}
//
//	res := make([]int, 0, size)
//
//	for tail != nil {
//		res = append(res, tail.Val.Val)
//		tail = tail.Prev
//	}
//
//	return res
//}
