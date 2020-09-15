package problems

/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	if n == 1 {
		return 1
	}

	start := 1
	end := n
	for start < end {
		mid := (start + end) / 2
		switch {
		case isBadVersion(mid):
			if !isBadVersion(mid - 1) {
				return mid
			}

			end = mid - 1
		case !isBadVersion(mid):
			start = mid + 1
		}
	}

	return end
}

// @lc code=end
