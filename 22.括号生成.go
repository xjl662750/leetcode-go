/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (72.17%)
 * Likes:    621
 * Dislikes: 0
 * Total Accepted:    54.6K
 * Total Submissions: 75.3K
 * Testcase Example:  '3'
 *
 * 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
 *
 * 例如，给出 n = 3，生成结果为：
 *
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 *
 *
 */

// @lc code=start
var ma map[string]int

func generateParenthesis(n int) []string {
	ma = make(map[string]int)
	generateSub("", n, n, 0)
	var result []string
	for k, _ := range ma {
		result = append(result, k)
	}
	return result
}

func generateSub(old string, m, n, a int) {
	if m == 0 && n == 1 {
		ma[old+")"] = 0
		return
	}
	if m == 0 {
		generateSub(old+")", m, n-1, a-1)
		return
	}
	if a == 0 {
		generateSub(old+"(", m-1, n, a+1)
		return
	}
	generateSub(old+"(", m-1, n, a+1)
	generateSub(old+")", m, n-1, a-1)
	return
}

// @lc code=end
