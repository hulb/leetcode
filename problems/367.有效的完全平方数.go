package problems

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	start := 1
	end := num
	for start < end {
		mid := (start + end) / 2
		switch {
		case mid*mid == num:
			return true
		case mid*mid > num:
			end = mid
		case mid*mid < num:
			start++
		}
	}

	return false
}

// @lc code=end
