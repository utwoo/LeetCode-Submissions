/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	if x < 10 {
		return true
	}

	y := 0
	for y < x {
		y = y*10 + x%10
		x = x / 10
	}

	return x == y || x == y/10
}
