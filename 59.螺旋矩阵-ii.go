/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	var matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		var arr = make([]int, n)
		matrix[i] = arr
	}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	i, j, a, b := 0, 0, 0, 1
	for n := 0; n < len(matrix)*len(matrix[0]); n++ {
		matrix[i][j] = n + 1
		if i+a < top {
			left = j + 1
			a = 0
			b = 1
		}
		if i+a > bottom {
			right = j - 1
			a = 0
			b = -1
		}
		if j+b < left {
			bottom = i - 1
			b = 0
			a = -1
		}
		if j+b > right {
			top = i + 1
			b = 0
			a = 1
		}
		i, j = i+a, j+b
	}
	return matrix
}

// @lc code=end

