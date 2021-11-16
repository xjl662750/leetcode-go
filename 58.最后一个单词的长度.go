/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	l := len(s)
	var num int
	var find bool
	for i := l - 1; i >= 0; i-- {
		if find && s[i] == ' ' {
			break
		}
		if s[i] >= 'A' && s[i] <= 'z' {
			find = true
			num++
		}
	}
	return num
}

// @lc code=end

