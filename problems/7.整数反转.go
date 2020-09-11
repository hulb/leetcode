package problems

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {

	nagative := x < 0
	if nagative {
		x = -1 * x
	}

	xstr := fmt.Sprintf("%d", x)
	length := len(xstr)

	arr := make([]uint8, length, length)
	for idx := range arr {
		if idx > 0 {
			realx := x
			for i := idx; i > 0; i-- {
				realx = realx - int(arr[idx-i])*(int(math.Pow10(length-(idx-i)-1)))
			}

			if idx < length-1 {
				arr[idx] = uint8(realx / (int(math.Pow10(length - idx - 1))))
			} else {
				arr[idx] = uint8(realx)
			}

			if idx == length-1 && arr[idx] == 0 {
				arr = arr[0 : length-1]
			}
		} else {
			arr[idx] = uint8(x / (int(math.Pow10(length - idx - 1))))
		}

		// fmt.Println(arr)
	}

	var reversex int
	for i := len(arr) - 1; i >= 0; i-- {
		reversex = reversex + int(arr[i])*int(math.Pow10(i))
	}

	if nagative {
		reversex = -1 * reversex
	}

	if float64(reversex) > math.Pow(2.0, 31.0)-1 || float64(reversex) < math.Pow(-2.0, 31.0) {
		return 0
	}

	return reversex
}

// @lc code=end
