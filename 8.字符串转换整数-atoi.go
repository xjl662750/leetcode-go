import (
	"math"
)

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	var isNegative, isOverFlow, start bool
	var res int
	for _, v := range str {
		if start && (v == ' ' || v == '-' || v == '+') {
			break
		}
		if v == '-' {
			isNegative = true
		}
		if v == '-' || v == '+' {
			if !start {
				start = true
			} else {
				break
			}
		}
		if v >= '0' && v <= '9' {
			if !start {
				start = true
			}
			res = res*10 + int(v-'0')
			result := 0
			result, isOverFlow = dealOverFlow(res, isNegative)
			if isOverFlow {
				res = result
				break
			}
		} else {
			if v != ' ' && v != '-' && v != '+' {
				break
			}
		}

	}
	if isNegative && !isOverFlow {
		res = -res
	}
	return res
}

func dealOverFlow(num int, isNegative bool) (res int, isOverFlow bool) {
	res = num
	if isNegative {
		res = -num
	}
	if res > math.MaxInt32 {
		isOverFlow = true
		res = math.MaxInt32
	}
	if res < math.MinInt32 {
		res = math.MinInt32
		isOverFlow = true
	}
	return res, isOverFlow

}

// @lc code=end
