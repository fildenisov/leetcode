package main

import "fmt"

func main() {

	n7 := ListNode{Val: 7}
	n6 := ListNode{Val: 6, Next: &n7}
	n5 := ListNode{Val: 5, Next: &n6}
	n4 := ListNode{Val: 4, Next: &n5}
	n3 := ListNode{Val: 3, Next: &n4}
	n2 := ListNode{Val: 2, Next: &n3}
	n1 := ListNode{Val: 1, Next: &n2}

	n7.Next = &n4

	fmt.Println(hasCycle(&n1))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type ReversedNode struct {
	Val  *ListNode
	Next *ReversedNode
}

func hasCycle(head *ListNode) bool {
	var tail, forCheck *ReversedNode
	for head != nil {
		forCheck = tail
		for forCheck != nil && forCheck.Next != nil {
			if forCheck.Next.Val == head {
				return true
			}
			forCheck = forCheck.Next
		}

		head = head.Next
		tail = &ReversedNode{
			Val:  head,
			Next: tail,
		}
	}
	return false
}
