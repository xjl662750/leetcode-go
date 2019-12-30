/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (58.95%)
 * Likes:    881
 * Dislikes: 0
 * Total Accepted:    97.1K
 * Total Submissions: 164.4K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 *
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 *
 *
 *
 * 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 *
 *
 * 示例:
 *
 * 输入: [1,8,6,2,5,4,8,3,7]
 * 输出: 49
 *
 */

// @lc code=start
//公用
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//方法一：暴力法
// 复杂度分析
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)

// func maxArea(height []int) int {
// 	var maxValue int
// 	for i := 0; i < len(height)-1; i++ {
// 		for j := len(height) - 1; i < j; j-- {
// 			maxValue = max(maxValue, (j-i)*min(height[i], height[j]))
// 		}

// 	}
// 	return maxValue
// }

//方法二：双指针法
func maxArea(height []int) int {
	var maxValue, i int
	var j = len(height) - 1
	for i < j {
		maxValue = max(maxValue, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxValue
}

// @lc code=end
