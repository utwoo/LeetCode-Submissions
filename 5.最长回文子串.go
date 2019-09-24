/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
func preProcess(s string) string {
	if len(s) == 0 {
		return "^$"
	}

	ret := "^"
	for _, val := range s {
		ret += "#" + string(val)
	}

	ret += "#$"
	return ret
}

func longestPalindrome(s string) string {
	T := preProcess(s)
	n := len(T)
	P := make([]int, n)
	center, right := 0, 0

	for i := 1; i < n-1; i++ {
		iMirror := 2*center - i // equals to i' =  center - (i - center)

		if right > i {
			if right-i > P[iMirror] {
				P[i] = P[iMirror]
				continue
			} else {
				P[i] = right - i
			}
		} else {
			P[i] = 0
		}

		// Attempt to expand palindrome centered at i
		for T[i+1+P[i]] == T[i-1-P[i]] {
			P[i]++
		}

		// If palindrome centered at i expand past right,
		// adjust center based on expanded palindrome
		if i+P[i] > right {
			center = i
			right = i + P[i]
		}
	}

	// find th maximun element in P
	maxLen := 0
	centerIndex := 0
	for i := 1; i < n-1; i++ {
		if P[i] > maxLen {
			maxLen = P[i]
			centerIndex = i
		}
	}

	index := (centerIndex - 1 - maxLen) / 2
	return s[index : index+maxLen]
}
