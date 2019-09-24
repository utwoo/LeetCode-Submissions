/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */
func isPowerOfTwo(n int) bool {
	return (n&(n-1) == 0 && n > 0)
}
