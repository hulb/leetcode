package problems

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	return describe(countAndSay(n - 1))
}

func describe(nstr string) string {
	count := -1
	var last rune
	var result string
	for idx, s := range nstr {
		if idx == 0 {
			count = 1
			last = s
			continue
		}

		if s != last {
			result += fmt.Sprintf("%d%s", count, string(last))
			count = 1
			last = s
		} else {
			count++
		}
	}

	result += fmt.Sprintf("%d%s", count, string(last))

	return result
}

// @lc code=end
