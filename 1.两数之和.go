/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
func twoSum(nums []int, target int) []int {
	dic := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if index, ok := dic[complement]; ok {
			return []int{index, i}
		}
		dic[nums[i]] = i
	}
	return nil
}
