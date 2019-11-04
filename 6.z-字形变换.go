/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	var arr []byte
	num := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; i+j < len(s); j = j + num {
			arr = append(arr, s[i+j])
			if i != 0 && i != numRows-1 && (num-i+j) < len(s) {
				arr = append(arr, s[num-i+j])
			}
		}
	}
	return string(arr)
}

// @lc code=end
