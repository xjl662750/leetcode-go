import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	var m string
	var str string
	var num int
	m = "1"
	for i := 1; i < n; i++ {
		str = ""
		num = 1
		for j := 0; j < len(m); j++ {
			if len(m) > j+1 && m[j] == m[j+1] {
				num++
			} else {
				temp := strconv.Itoa(num)
				str = str + temp + string(m[j])
				num = 1
			}
		}
		m = str
	}
	fmt.Println(m)
	return m
}

// @lc code=end

