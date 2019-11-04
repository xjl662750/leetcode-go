/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var y int
	for i := x; i != 0; i = i / 10 {
		y = y*10 + i%10
	}
	if x == y {
		return true
	}
	return false
}

// @lc code=end
