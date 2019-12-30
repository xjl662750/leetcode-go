/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (51.29%)
 * Likes:    493
 * Dislikes: 0
 * Total Accepted:    57.5K
 * Total Submissions: 111.7K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */

// @lc code=start
var m map[string][]string

var res []string

func letterCombinations(digits string) []string {
	res = []string{}
	m = make(map[string][]string)
	m["2"] = []string{"a", "b", "c"}
	m["3"] = []string{"d", "e", "f"}
	m["4"] = []string{"g", "h", "i"}
	m["5"] = []string{"j", "k", "l"}
	m["6"] = []string{"m", "n", "o"}
	m["7"] = []string{"p", "q", "r", "s"}
	m["8"] = []string{"t", "u", "v"}
	m["9"] = []string{"w", "x", "y", "z"}
	if len(digits) > 0 {
		sub(digits, "")
	}
	return res
}
func sub(digits string, s string) {
	length := len(digits)
	for _, v := range m[string(digits[0])] {
		t := s + v
		if length > 1 {
			sub(digits[1:], t)
		} else {
			res = append(res, t)
		}
	}
}

// @lc code=end
