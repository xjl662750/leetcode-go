/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	var m = make(map[int]int)
	m[1] = 1
	m[2] = 2
	for i := 3; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
	}

	return m[n]
}

// @lc code=end

