import "container/list"

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
func isValid(s string) bool {
	symbolList := list.New()
	symbolMap := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := 0; i < len(s); i++ {
		val := s[i]
		switch val {
		case '[', '{', '(':
			symbolList.PushBack(val)
		case ']', '}', ')':
			if symbolList.Len() == 0 {
				return false
			}

			last := symbolList.Back()
			if last.Value != symbolMap[val] {
				return false
			}
			symbolList.Remove(last)
		}
	}

	return symbolList.Len() == 0
}
