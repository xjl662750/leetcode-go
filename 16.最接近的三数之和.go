import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode-cn.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (41.93%)
 * Likes:    291
 * Dislikes: 0
 * Total Accepted:    55.1K
 * Total Submissions: 130.9K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target
 * 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 *
 * 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
 *
 * 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 *
 *
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	// res := nums[0] + nums[1] + nums[2]
	// min := abs(res - target)
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		for k := j + 1; k < len(nums); k++ {
	// 			temp := nums[i] + nums[j] + nums[k]
	// 			if abs(temp-target) < min {
	// 				min = abs(temp - target)
	// 				res = temp
	// 			}
	// 		}
	// 	}
	// }
	// return res
	sort.Ints(nums)
	length := len(nums)
	res := nums[0] + nums[1] + nums[2]
	min := abs(res - target)
	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := length - 1
		for j := i + 1; j < k; {
			temp := nums[i] + nums[j] + nums[k]
			if temp == target {
				return target
			}
			if temp < target {
				j++
			}
			if temp > target {
				k--
			}
			if abs(temp-target) < min {
				res = temp
				min = abs(temp - target)
			}

		}
	}
	return res
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// @lc code=end
