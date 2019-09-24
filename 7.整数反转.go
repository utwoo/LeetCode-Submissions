import "math"

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
func reverse(x int) int {
	var result int
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return result
}
