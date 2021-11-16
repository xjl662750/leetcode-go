/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	var newIntervals [][]int
	var m int
	for k, v := range intervals {
		m = k + 1
		if v[0] > newInterval[1] {
			m = k
			break
		}
		if v[1] < newInterval[0] {
			newIntervals = append(newIntervals, v)
			continue
		}
		if v[0] < newInterval[0] && v[1] >= newInterval[0] {
			newInterval[0] = v[0]
		}
		if v[0] <= newInterval[1] && v[1] >= newInterval[1] {
			newInterval[1] = v[1]
			break
		}

	}
	newIntervals = append(newIntervals, newInterval)
	newIntervals = append(newIntervals, intervals[m:]...)
	return newIntervals
}

// @lc code=end

