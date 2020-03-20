/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode-cn.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (59.38%)
 * Likes:    221
 * Dislikes: 0
 * Total Accepted:    45.8K
 * Total Submissions: 75.3K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用一次。
 *
 * 说明：
 *
 *
 * 所有数字（包括目标数）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 所求解集为:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 所求解集为:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	quickSort(candidates, 0, len(candidates)-1)
	res := getTemp(candidates, target, -1)
	return res
}

func quickSort(arr []int, left, right int) {
	if left < right {
		index := getIndex(arr, left, right)
		quickSort(arr, left, index-1)
		quickSort(arr, index+1, right)
	}

}
func getIndex(arr []int, left, right int) int {
	temp := arr[left]
	for left < right {
		for left < right && temp <= arr[right] {
			right--
		}
		swap(arr, left, right)
		for left < right && temp >= arr[left] {
			left++
		}
		swap(arr, left, right)
	}
	return left
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func getTemp(candidates []int, target int, ck int) (ret [][]int) {
	if target <= 0 {
		return
	}

	for i := ck + 1; i < len(candidates); i++ {
		if target == candidates[i] {
			ret = append(ret, [][]int{[]int{candidates[i]}}...)
			break
		}
		t := getTemp(candidates, target-candidates[i], i)
		if t != nil {
			for k, _ := range t {
				t[k] = append([]int{candidates[i]}, t[k]...)
			}
			ret = append(ret, t...)
		}
		for i < len(candidates)-1 && candidates[i+1] == candidates[i] {
			i++
		}
	}
	return
}

// @lc code=end

