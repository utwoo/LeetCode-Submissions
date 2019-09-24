import "strconv"

/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3的幂
 */
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	x := strconv.FormatInt(int64(n), 3)
	y := strconv.FormatInt(int64(n-1), 3)

	return len(x) == len(y)+1
}
