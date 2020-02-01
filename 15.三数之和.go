/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
    var result [][]int

	sort.Ints(nums)

	if (len(nums) < 2) || (nums[0] > 0) || (nums[len(nums)-1] < 0) {
		return result
	}

	for k := 0; k < len(nums)-2; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		if nums[k] > 0 {
			break
		}
		
		i := k + 1
		j := len(nums) - 1

		for i < j && nums[j] >= 0 {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[k], nums[i], nums[j]})

				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}

				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else if nums[i]+nums[j]+nums[k] < 0 {
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			} else {
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}
		}
	}
	return result
}
// @lc code=end

