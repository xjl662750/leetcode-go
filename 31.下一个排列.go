/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 *
 * https://leetcode-cn.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (32.17%)
 * Likes:    332
 * Dislikes: 0
 * Total Accepted:    33.5K
 * Total Submissions: 103.7K
 * Testcase Example:  '[1,2,3]'
 *
 * 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
 * 
 * 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
 * 
 * 必须原地修改，只允许使用额外常数空间。
 * 
 * 以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 * 
 */

// @lc code=start
func nextPermutation(nums []int)  {
    for i:=len(nums)-1;i>=0;i--{
		if i==0{
			swapSort(nums)
			break
		}
		if i>0&&nums[i]>nums[i-1]{
			for m:=len(nums)-1;m>=i;m--{
				if nums[m]>nums[i-1]{
					nums[i-1],nums[m]=nums[m],nums[i-1]
					break
				}
			}
			swapSort(nums[i:])
			break
		}
	}
}
func swapSort(nums []int){
	for j,k:=0,len(nums)-1;j<k;j++{
		nums[j],nums[k]=nums[k],nums[j]
		k--
	}
}
// @lc code=end

