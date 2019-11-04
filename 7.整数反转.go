import (
	"math"
)

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	// MaxInt32 := 1<<31 - 1
	// MinInt32 := -1 << 31
	var res int
	for i := x; i != 0; i = i / 10 {
		res = res*10 + i%10
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}

// @lc code=end
