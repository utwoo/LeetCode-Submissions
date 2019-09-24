import "strings"

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */
func intToRoman(num int) string {
	result := ""

	num_1000 := num / 1000
	num_100 := (num % 1000) / 100
	num_10 := (num % 100) / 10
	num_1 := (num % 10)

	result += strings.Repeat("M", num_1000)

	if num_100 == 9 {
		result += "CM"
	} else if num_100 == 4 {
		result += "CD"
	} else {
		result += strings.Repeat("D", num_100/5) + strings.Repeat("C", num_100%5)
	}

	if num_10 == 9 {
		result += "XC"
	} else if num_10 == 4 {
		result += "XL"
	} else {
		result += strings.Repeat("L", num_10/5) + strings.Repeat("X", num_10%5)
	}

	if num_1 == 9 {
		result += "IX"
	} else if num_1 == 4 {
		result += "IV"
	} else {
		result += strings.Repeat("V", num_1/5) + strings.Repeat("I", num_1%5)
	}

	return result
}
