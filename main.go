package main

import (
	"fmt"
	"leetcode/problems"
)

func main() {
	fmt.Println("hello world")

	// l1 := &problems.ListNode{
	// 	Val: 1,
	// 	Next: &problems.ListNode{
	// 		Val: 2,
	// 		Next: &problems.ListNode{
	// 			Val: 4,
	// 		},
	// 	},
	// }
	// l2 := &problems.ListNode{
	// 	Val: 1,
	// 	Next: &problems.ListNode{
	// 		Val: 3,
	// 		Next: &problems.ListNode{
	// 			Val: 4,
	// 		},
	// 	},
	// }

	// problems.MergeTwoLists(l1, l2).Println()
	// in := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// fmt.Println(problems.RemoveDuplicates(in))
	// fmt.Println(in)

	// in := []int{3, 3}
	// fmt.Println(problems.RemoveElement(in, 3))
	// fmt.Println(in)

	// in := []int{1, 2, 3, 4, 5, 10}
	// fmt.Println(problems.SearchInsert(in, 2))

	// fmt.Println(problems.RomanToInt("MCMXCIV"))

	// fmt.Println(problems.CountAndSay(11))

	// fmt.Println(problems.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	// fmt.Println(problems.MaxSubArray([]int{-2, -1}))

	fmt.Println(problems.MySqrt(8))
	fmt.Println(problems.MySqrt(9))
	fmt.Println(problems.MySqrt(10))
}
