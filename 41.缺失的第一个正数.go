/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 *
 * https://leetcode-cn.com/problems/first-missing-positive/description/
 *
 * algorithms
 * Hard (37.32%)
 * Likes:    423
 * Dislikes: 0
 * Total Accepted:    41.6K
 * Total Submissions: 110.3K
 * Testcase Example:  '[1,2,0]'
 *
 * 给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
 *
 * 示例 1:
 *
 * 输入: [1,2,0]
 * 输出: 3
 *
 *
 * 示例 2:
 *
 * 输入: [3,4,-1,1]
 * 输出: 2
 *
 *
 * 示例 3:
 *
 * 输入: [7,8,9,11,12]
 * 输出: 1
 *
 *
 * 说明:
 *
 * 你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。
 *
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	l := len(nums)
	s := make([]int, l+2)
	for _, v := range nums {
		if v > 0 && v <= l+1 {
			s[v] = 1
		}
	}
	for i := 1; i <= l+1; i++ {
		if s[i] == 0 {
			return i
		}
	}
	return 0
}

// func firstMissingPositive(nums []int) int {
// 	var m = make(map[int]bool)
// 	for i := 0; i < len(nums); i++ {
// 		m[nums[i]] = true
// 	}
// 	var t int = 1
// 	for true {
// 		if m[t] {
// 			t++
// 			continue
// 		}
// 		break
// 	}

// 	return t
// }

// func firstMissingPositive(nums []int) int {
// 	var flag = true
// 	var t int = 1
// 	for flag {
// 		i := 0
// 		for ; i < len(nums); i++ {
// 			if nums[i] == t {
// 				t++
// 				break
// 			}
// 		}
// 		if i == len(nums) {
// 			flag = false
// 		}
// 	}
// 	return t
// }

// @lc code=end

