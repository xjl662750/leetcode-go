/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	result := [][]int{}
	sort.Ints(nums)
	helper(nums, 0, &result)

	return result
}

// 回溯函数实现
// i表示本次函数需要放置的元素位置
func helper(nums []int, i int, result *[][]int) {
	n := len(nums)
	if i == n-1 {
		tmp := make([]int, n)
		copy(tmp, nums)
		*result = append(*result, tmp)
		return
	}
	// nums[0:i]是已经决定的部分，nums[i:]是待决定部分，同时待选元素也都在nums[i:]
	for k := i; k < n; k++ {
		// 跳过重复数字
		if k != i && nums[k] == nums[i] {
			continue
		}
		nums[k], nums[i] = nums[i], nums[k]
		helper(nums, i+1, result)
	}
	// 还原状态
	for k := n - 1; k > i; k-- {
		nums[i], nums[k] = nums[k], nums[i]
	}
}


// @lc code=end

