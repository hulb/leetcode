package problems

import "fmt"

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	l1n1 := l1.Val
	l2n1 := l2.Val

	start1 := l1.Next
	start2 := l2

	merge := ListNode{Val: l1n1}
	if l1n1 > l2n1 {
		merge.Val = l2n1
		start1 = l2.Next
		start2 = l1
	}

	mergeP := &merge
	for start1 != nil && start2 != nil {
		if start1.Val > start2.Val {
			mergeP.Next = &ListNode{Val: start2.Val}
			start2 = start2.Next
		} else {
			mergeP.Next = &ListNode{Val: start1.Val}
			start1 = start1.Next
		}

		mergeP = mergeP.Next
	}

	if start1 != nil {
		mergeP.Next = start1
	} else if start2 != nil {
		mergeP.Next = start2
	}

	return &merge
}

func (l ListNode) Println() {
	fmt.Print(l.Val)
	t := &l
	for t.Next != nil {
		fmt.Print("->")
		t = t.Next
		fmt.Print(t.Val)
	}
	fmt.Println()
}

// @lc code=end
