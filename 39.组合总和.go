/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := getTemp(candidates, target, 0)
	return res
}

func getTemp(candidates []int, target int, ck int) (ret [][]int) {
	if target <= 0 {
		return
	}

	for i := ck; i < len(candidates); i++ {
		if target == candidates[i] {
			ret = append(ret, [][]int{[]int{candidates[i]}}...)
		}
		t := getTemp(candidates, target-candidates[i], i)
		if t != nil {
			for k, _ := range t {
				t[k] = append([]int{candidates[i]}, t[k]...)
			}
			ret = append(ret, t...)
		}
	}
	return
}

// @lc code=end

