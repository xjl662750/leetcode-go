/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	var result []int
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	i, j, a, b := 0, 0, 0, 1
	for n := 0; n < len(matrix)*len(matrix[0]); n++ {
		result = append(result, matrix[i][j])
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
	return result
}

// @lc code=end

