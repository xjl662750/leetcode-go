import "fmt"
/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 *
 * https://leetcode-cn.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (18.73%)
 * Likes:    229
 * Dislikes: 0
 * Total Accepted:    28.3K
 * Total Submissions: 149.6K
 * Testcase Example:  '10\n3'
 *
 * 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
 * 
 * 返回被除数 dividend 除以除数 divisor 得到的商。
 * 
 * 示例 1:
 * 
 * 输入: dividend = 10, divisor = 3
 * 输出: 3
 * 
 * 示例 2:
 * 
 * 输入: dividend = 7, divisor = -3
 * 输出: -2
 * 
 * 说明:
 * 
 * 
 * 被除数和除数均为 32 位有符号整数。
 * 除数不为 0。
 * 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。
 * 
 * 
 */

// @lc code=start
// func divide(dividend int, divisor int) int {
// 	if divisor==-1&&dividend<=(-1<<31){
// 		dividend=1<<31-1
// 		return dividend
// 	}
// 	if divisor==1{
// 		return dividend
// 	}
// 	t:=1
// 	if divisor<0{
// 			divisor=-divisor
// 			t=-t
// 	}
// 	if dividend<0{
// 			dividend=-dividend
// 			t=-t
// 	}
	
// 	if t<0{
// 		return -divideSub(dividend,divisor)
// 	}
//     return divideSub(dividend,divisor)
// }
// func divideSub(dividend int, divisor int) int {
// 	if dividend<divisor{
// 		return 0
// 	}
// 	temp:=divisor
// 	var res int
// 	res++
// 	for dividend>temp+temp{
// 		res=res+res
// 		temp=temp+temp
// 	}
	
//     return divideSub(dividend-temp,divisor)+res
// }
func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}

	if divisor == 1 {
		return dividend
	}

	maxInt := 1<<31

	if dividend == maxInt-1 && divisor == -1 {
		return maxInt-1
	}
	isMinus := (dividend < 0 && divisor >0 )|| (dividend > 0 && divisor < 0 )
	var (
		ds =  divisor
		dd = dividend
	)
	
	if ds < 0 {
		ds = -ds
	}

	if dd < 0 {
		dd = -dd
	}

	var result int
	for i := 31; i >= 0; i-- {
        it := uint(i)
		if dd >> it >= ds {
			result += 1<<it
			dd -= ds<<it
		}
	}

	if isMinus {
		result = -result
	}

	if result > maxInt-1 {
		result = maxInt-1
	}

	return result
}


// @lc code=end

