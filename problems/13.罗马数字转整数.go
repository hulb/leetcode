package problems

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	mapping := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	sArr := []rune(s)
	result := 0
	for i := 0; i < len(sArr); i++ {
		v, ok := mapping[string(sArr[i])]
		if ok {
			result += v
		}

		if i < len(sArr)-1 {
			if v2, exist := mapping[string(sArr[i:i+2])]; exist {
				result = result - v + v2
				i++
			}
		}
	}

	return result
}

// @lc code=end
