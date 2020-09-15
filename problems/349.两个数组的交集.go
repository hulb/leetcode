package problems

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]struct{})
	for _, i := range nums1 {
		nums1Map[i] = struct{}{}
	}

	intersec := make(map[int]struct{})
	for _, i := range nums2 {
		if _, ok := nums1Map[i]; ok {
			intersec[i] = struct{}{}
		}
	}

	var result []int
	for i := range intersec {
		result = append(result, i)
	}

	return result
}

// @lc code=end
