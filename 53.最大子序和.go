/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */
func maxSubArray(nums []int) int {
	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX

	result := INT_MIN

	sum := 0

	for _, val := range nums {
		sum = max(val, sum+val)
		result = max(sum, result)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
