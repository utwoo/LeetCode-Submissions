/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
func lengthOfLongestSubstring(s string) int {
	positions := make(map[rune]int)
	former, result := 0, 0

	for current, value := range s {
		n, ok := positions[value]
		if ok {
			if n > former {
				former = n
			}
		}

		if current-former+1 > result {
			result = current - former + 1
		}

		positions[value] = current + 1
	}

	return result
}
