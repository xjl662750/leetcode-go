import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (24.31%)
 * Likes:    1519
 * Dislikes: 0
 * Total Accepted:    117.8K
 * Total Submissions: 482.2K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 * 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */

// @lc code=start
func threeSum(num []int) [][]int {
	sort.Ints(num)
	res := [][]int{}
	length := len(num)
	for i := 0; i < length; i++ {
		if num[i] > 0 {
			continue //当前>0 则不用计算之后的
		}
		if i > 0 && num[i] == num[i-1] {
			continue //当前重复的元素,之前计算过一次，不必再次计算
		}
		z := length - 1
		for j := i + 1; j < length && z > i && j < z; {
			if z < length-1 && num[z] == num[z+1] {
				z-- //去重
				continue
			}
			if num[j] == num[j-1] && j-1 > i {
				j++ //去重
				continue
			}
			if num[i]+num[j]+num[z] > 0 {
				z--
			} else if num[i]+num[j]+num[z] < 0 {
				j++
			} else {
				arr := []int{num[i], num[j], num[z]}
				res = append(res, arr)
				z--
				j++
			}
		}
	}
	return res
}

// @lc code=end
