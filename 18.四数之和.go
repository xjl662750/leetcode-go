import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode-cn.com/problems/4sum/description/
 *
 * algorithms
 * Medium (36.02%)
 * Likes:    318
 * Dislikes: 0
 * Total Accepted:    40.5K
 * Total Submissions: 112.5K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c
 * + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
 *
 * 注意：
 *
 * 答案中不可以包含重复的四元组。
 *
 * 示例：
 *
 * 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
 *
 * 满足要求的四元组集合为：
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 *
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	length := len(nums)
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for n := length - 1; n > i+2; n-- {
			if n < length-1 && nums[n] == nums[n+1] {
				continue
			}
			m := n - 1
			for j := i + 1; j < m; j++ {
				if j > i+1 && nums[j] == nums[j-1] {
					continue
				}
				temp := nums[i] + nums[j] + nums[m] + nums[n]
				if temp > target {
					m--
					j--
				}
				if temp == target {
					res = append(res, []int{nums[i], nums[j], nums[m], nums[n]})
				}

			}
		}

	}
	return res
}

// @lc code=end
