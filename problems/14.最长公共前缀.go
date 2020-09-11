package problems

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	minLength := len(strs[0])
	for _, s := range strs[1:] {
		if len(s) < minLength {
			minLength = len(s)
		}
	}

	var common []byte
	for i := 0; i <= minLength-1; i++ {
		var lastStr byte
		commonExist := false
		for idx, s := range strs {
			if idx == 0 {
				lastStr = s[i]
				continue
			}

			if s[i] != lastStr {
				commonExist = false
				break
			} else {
				commonExist = true
			}
		}

		if commonExist {
			common = append(common, lastStr)
		} else {
			break
		}
	}

	return string(common)
}

// @lc code=end
