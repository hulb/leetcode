package problems

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	start := 1
	end := x
	for start < end {
		mid := (start + end) / 2
		switch {
		case mid*mid == x:
			return mid
		case mid*mid > x:
			end = mid
		case mid*mid < x:
			start++
		}
	}

	return start - 1
}

// @lc code=end
