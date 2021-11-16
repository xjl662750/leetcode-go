/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	var m = make(map[int]int)
	l := len(nums)
	for i := 0; i < l; i++ {
		m[nums[i]]++
		if m[nums[i]] > l/2 {
			return nums[i]
		}
	}
	return 0
}

// @lc code=end

