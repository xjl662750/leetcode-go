/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	start, end := 0, 0
	var zero bool
	for _, v := range intervals {
		if v[0] == 0 {
			zero = true
		}
		if v[0] > end {
			if end != 0 || zero {
				res = append(res, []int{start, end})
			}
			start = v[0]
		}
		if v[1] > end {
			end = v[1]
		}
	}
	res = append(res, []int{start, end})
	return res
}

// @lc code=end

