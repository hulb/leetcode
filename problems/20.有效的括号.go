package problems

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start

var mapping = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	arr := []byte(s)

	if len(arr)%2 != 0 {
		return false
	}

	stack := make([]byte, 0, len(arr)/2)
	for _, i := range arr {
		if _, ok := mapping[i]; !ok {
			if len(stack) == 0 || mapping[stack[len(stack)-1]] != i {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, i)
		}
	}

	return len(stack) == 0
}

// @lc code=end
