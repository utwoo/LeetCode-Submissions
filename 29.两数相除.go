/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	// int32 max value
	var INT_MAX = int32(^uint32(0) >> 1)
	// int32 min value
	var INT_MIN = int32(^INT_MAX)
	// check overflow
	if dividend == int(INT_MIN) && divisor == -1 {
		return int(INT_MAX)
	}
	// check sign
	flag := false
	if dividend > 0 && divisor > 0 {
		flag = true
	} else if dividend < 0 && divisor < 0 {
		flag = true
	}
	
	n1 := int(math.Abs(float64(dividend)))
	n2 := int(math.Abs(float64(divisor)))

	// divide by bit 
	result := 0
	for n2 <= n1 {
		shift := 0
		for n1 >= (n2 << shift) {
			shift++
		}
		n1 = n1 - (n2 << (shift - 1))
		result = result + (1 << (shift - 1))
	}

	if !flag {
		result = result * -1
	}

	return result
}
// @lc code=end

