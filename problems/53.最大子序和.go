package problems

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
// unexceptable
func maxSubArray2(nums []int) int {
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			if sum(nums[i:j]) > max {
				max = sum(nums[i:j])
			}
		}
	}

	return max
}

func sum(nums []int) int {
	var result int
	for _, i := range nums {
		result += i
	}

	return result
}

// 这里存在一个最优子结构,最优子结构的意思是局部最优解能决定全局最优解
func maxSubArray(nums []int) int {
	max := nums[0]
	sum := nums[0]
	for _, i := range nums[1:] {
		if sum+i > i { // 比较i和i之前所有数之合，去较大的那一个
			sum += i
		} else {
			sum = i
		}

		if sum > max {
			max = sum
		}
	}

	return max
}

// @lc code=end
