/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
// func canJump(nums []int) bool {
// 	var max int
// 	for i := 0; i < len(nums); i++ {
// 		if i <= max {
// 			if nums[i]+i > max {
// 				max = nums[i] + i
// 			}
// 			if max >= len(nums)-1 {
// 				return true
// 			}
// 		}
// 		if i > max {
// 			return false
// 		}
// 	}
// 	return false
// }

// func canJump(nums []int) bool {
// 	var max int
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i]+i > max {
// 			max = nums[i] + i
// 		}
// 		if i >= max && i < len(nums)-1 {
// 			return false
// 		}
// 	}
// 	return true
// }
func canJump(nums []int) bool {
	target := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= target {
			target = i
		}
	}
	return target == 0
}

// @lc code=end

