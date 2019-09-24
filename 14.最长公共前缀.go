import "sort"

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)

	if strs[0] == "" {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	first := strs[0]
	last := strs[len(strs)-1]

	i := 0
	for ; i < len(first); i++ {
		if i >= len(last) || first[i] != last[i] {
			return first[:i]
		}
	}
	return first[:i]
}
