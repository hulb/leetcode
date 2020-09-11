package problems

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xstrArr := []rune(fmt.Sprintf("%d", x))
	length := len(xstrArr)

	for i := 0; i <= length/2; i++ {
		if xstrArr[i] != xstrArr[length-i-1] {
			return false
		}
	}

	return true
}

// @lc code=end
