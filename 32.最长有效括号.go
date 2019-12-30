/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 *
 * https://leetcode-cn.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (28.52%)
 * Likes:    430
 * Dislikes: 0
 * Total Accepted:    31.2K
 * Total Submissions: 108K
 * Testcase Example:  '"(()"'
 *
 * 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
 * 
 * 示例 1:
 * 
 * 输入: "(()"
 * 输出: 2
 * 解释: 最长有效括号子串为 "()"
 * 
 * 
 * 示例 2:
 * 
 * 输入: ")()())"
 * 输出: 4
 * 解释: 最长有效括号子串为 "()()"
 * 
 * 
 */

// @lc code=start
// func longestValidParentheses(s string) int {
// 	var left,right,max int
// 	for i:=0;i<len(s);i++{
// 		if string(s[i])=="("{
// 			left++
// 		}else{
// 			right++
// 		}
// 		if left<right{
// 			left,right=0,0
// 		}
// 		if left==right{
// 			max=Max(max,2*left)
// 		}
// 	}
// 	left,right=0,0
// 	for i:=len(s)-1;i>=0;i--{
// 		if string(s[i])=="("{
// 			left++
// 		}else{
// 			right++
// 		}
// 		if left>right{
// 			left,right=0,0
// 		}
// 		if left==right{
// 			max=Max(max,2*left)
// 		}
// 	}
// 	return max
// }
//动态规划
func longestValidParentheses(s string) int {
	var dp=make([]int,len(s))
	var max int
	for i:=1;i<len(s);i++{
		if string(s[i])==")"{
			if string(s[i-1])=="("{
				if i>2{
					dp[i]=dp[i-2]+2
				}else{
					dp[i]=2
				}
			}else if i-dp[i-1]>0&& string(s[i-dp[i-1]-1])=="("{
				if i - dp[i - 1] >= 2{
					dp[i] = dp[i - 1]+ dp[i - dp[i - 1] - 2]  + 2
				}else{
					dp[i] =dp[i - 1]+2;
				}
			}
			max =Max(max, dp[i]);
		}
	}
	
	return max
}

func Max(x,y int)int{
	if x>y{
		return x
	}
	return y
}
// @lc code=end

