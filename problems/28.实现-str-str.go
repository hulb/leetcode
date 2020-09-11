package problems

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr2(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	index := -1
	needles := []rune(needle)
	stacks := []rune(haystack)
	for idx, s := range stacks {
		if s == needles[0] {
			index = idx
			idx++

			if idx >= len(stacks) && len(needles) > 1 {
				index = -1
				continue
			}

			for i := 1; i < len(needles); i++ {
				if stacks[idx] != needles[i] {
					index = -1
					break
				}

				if idx >= len(stacks)-1 && i < len(needles)-1 {
					index = -1
					break
				}

				idx++

			}

			if index != -1 {
				return index
			}
		}
	}

	return index
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	index := -1
	nIdx := 0
	needles := []rune(needle)
	stacks := []rune(haystack)

	for i := 0; i < len(stacks); i++ {
		if stacks[i] == needles[nIdx] {
			if index == -1 {
				index = i
			}

			nIdx++
			if nIdx > len(needles)-1 {
				break
			}

		} else {
			nIdx = 0
			if index != -1 {
				i = index
			}
			index = -1
		}
	}

	if index != -1 && nIdx <= len(needles)-1 {
		return -1
	}

	return index
}

// @lc code=end
