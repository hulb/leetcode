package problems

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	j := len(nums) - 1
	i := 0

	if (j == i && nums[i] == val) || j < 0 {
		return 0
	}

	for ; i < j && j >= 0 && i < len(nums); i++ {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
			i--
		}
	}

	if j == i && nums[i] == val {
		return i
	}

	return i + 1
}

// @lc code=end
