/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */

// @lc code=start
func isMatch(s, p string) bool {
	if s == "" {
		return p == "" || p == "*"
	}
	if p == "" {
		return false
	}
	sLen, pLen := len(s), len(p)

	match := make([][]bool, sLen+1)
	for i := 0; i < len(match); i++ {
		match[i] = make([]bool, pLen+1)
	}
	match[0][0] = true
	for i := 1; i < pLen+1; i++ {
		match[0][i] = match[0][i-1] && p[i-1] == '*'
	}
	for si := 1; si < sLen+1; si++ {
		for pi := 1; pi < pLen+1; pi++ {
			if p[pi-1] == '*' {
				match[si][pi] = match[si-1][pi] || match[si][pi-1]
			} else {
				if s[si-1] == p[pi-1] || p[pi-1] == '?' {
					match[si][pi] = match[si-1][pi-1]
				}
			}
		}
	}
	return match[sLen][pLen]
	
}

// @lc code=end

