package problems

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	last := nums[0]

	for i := 1; i < length; i++ {
		if nums[i] == last {
			if i+1 == length {
				nums = nums[0:i]
				break
			}

			nums = append(nums[0:i], nums[i+1:]...)
			length--
			i--
			continue
		}

		last = nums[i]
	}

	return len(nums)
}

func removeDuplicates(nums []int) int {

	if len(nums) < 2 {
		return 1
	}

	i, j := 0, 1
	for ; j < len(nums); j++ {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
	}

	return i + 1
}

// @lc code=end
