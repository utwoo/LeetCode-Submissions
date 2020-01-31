/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	// int32 max value
	var int32Max = (1 << 31) - 1
	// int32 min value
	var int32Min = int(^int32(int32Max))
	// check overflows
	if dividend == int32Min && divisor == -1 {
		return int(int32Max)
	}
	// check sign
	flag := (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0)

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
		result = -result
	}

	return result
}
// @lc code=end

