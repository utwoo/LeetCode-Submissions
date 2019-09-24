/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
func maxArea(height []int) int {
	area := 0
	left, right, max := 0, len(height)-1, 0
	for left < right {
		if height[left] > height[right] {
			area = height[right] * (right - left)
			right--
		} else {
			area = height[left] * (right - left)
			left++
		}

		if area > max {
			max = area
		}
	}
	return max
}
