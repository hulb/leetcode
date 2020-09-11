package problems

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert2(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}

	for idx, i := range nums {
		if i == target {
			return idx
		}

		if i < target {
			if idx != len(nums)-1 {
				if nums[idx+1] >= target {
					return idx + 1
				}
			} else {
				return len(nums)
			}
		}
	}

	return -1
}

func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}

	if target == nums[len(nums)-1] {
		return len(nums) - 1
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	start, end := 0, len(nums)-1
	for start < end {
		half := (start + end) / 2
		switch {
		case nums[half] == target:
			return half
		case nums[half] > target:
			end = half
		case nums[half] < target:
			start++
		}
	}

	return start
}

// @lc code=end
