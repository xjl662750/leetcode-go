/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

func lengthOfLongestSubstring(s string) int {
m,max,left := make(map[rune]int),0,0
    for k,v:=range s{
		if a,ok:=m[v];ok&&a>=left{
			if k-left>max{
			max=k-left
			}
			left=m[v]+1
		}
		m[v]=k
	}
	if len(s)-left > max {
		max = len(s) - left
	}
	return max
}

